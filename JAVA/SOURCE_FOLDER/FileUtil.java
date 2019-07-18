package util;

import java.io.BufferedInputStream;
import java.io.File;
import java.io.FileInputStream;
import java.io.InputStream;
import java.io.OutputStream;
import java.util.zip.ZipEntry;
import java.util.zip.ZipOutputStream;

import javax.servlet.http.HttpServletRequest;
import javax.servlet.http.HttpServletResponse;

public class FileUtil {
	
	public void fileDownload(HttpServletRequest request, HttpServletResponse response, String paramFileList, String fileName) throws Exception {
		request.setCharacterEncoding("UTF-8");
		
		int bufferSize = 1024*2;
		
		String[] fileList = paramFileList.split(",");
		
		InputStream is = null;
		OutputStream os = null;
		String client = request.getHeader("User-Agent");
		
		response.reset();
		response.setContentType("application/octet-stream");
		response.setHeader("Content-Transfer-Encoding", "binary");
		
		if(client.indexOf("MSIE")!=-1) {
			response.setHeader("Content-Disposition", "attachment; fileName=\""
					+ java.net.URLEncoder.encode(fileName + ((fileList.length>1) ? ".zip":""), "UTF-8").replaceAll("\\+", "\\ ") + "\"");
		}else if(client.indexOf("Trident")!=-1) {
			response.setHeader("Content-Disposition", "attachment; fileName=\""
					+ java.net.URLEncoder.encode(fileName + ((fileList.length>1) ? ".zip":""), "UTF-8").replaceAll("\\+", "\\ ") + "\"");
		}else {
			response.setHeader("Content-Disposition", "attachment; fileName=\""
					+ new String((fileName + ((fileList.length>1) ? ".zip":"" )).getBytes("UTF-8"),"8859_1")+"\"");
		}
		
		os = response.getOutputStream();
		
		if(fileList.length>1) {
			ZipOutputStream zos = new ZipOutputStream(os);
			zos.setLevel(8);
			BufferedInputStream bis = null;
			
			for(int i=0; i<fileList.length; i++) {
				File file = new File(fileList[i]);
				bis = new BufferedInputStream(new FileInputStream(file));
				ZipEntry zTry = new ZipEntry(fileList[i].substring(fileList[i].lastIndexOf("/")));
				zTry.setTime(file.lastModified());
				zos.putNextEntry(zTry);
				
				byte[] buffer = new byte[bufferSize];
				int cnt = 0;
				while((cnt = bis.read(buffer, 0, bufferSize)) != -1) {
					zos.write(buffer, 0, cnt);
				}
				zos.closeEntry();
				
				if(zos!=null) zos.close();
				if(bis!=null) bis.close();
				if(os!=null) os.close();
				
			}
		}else {
			File file = new File(fileList[0]);
			is = new FileInputStream(file);
			
			response.setHeader("Content-Length", ""+file.length());
			
			byte[] buffer = new byte[(int) file.length()];
			int cnt = 0;
			while((cnt = is.read(buffer)) != -1) {
				os.write(buffer, 0, cnt);
			}
			
			if(is!=null) is.close();
			if(os!=null) os.close();
		}
	}
	
}
