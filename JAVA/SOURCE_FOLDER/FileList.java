package com.test;

import java.io.File;
import java.io.FileFilter;
import java.io.FilenameFilter;
import java.util.ArrayList;
import java.util.Collections;
import java.util.Comparator;
import java.util.HashMap;
import java.util.List;
import java.util.Map;

public class FileList {
	public List<HashMap<String, String>> getFileList(Map param){
		List<HashMap<String, String>> allFileList = new ArrayList<HashMap<String, String>>();
		
		FileFilterUtil fileFilterUtil = new FileFilterUtil("report");
		FilenameFilterUtil filenameFilterUtil = new FilenameFilterUtil("report");
		
		//파일 날짜 순 정렬
		Comparator cprt = new Comparator<Map<String, String>>() {

			@Override
			public int compare(Map<String, String> map1, Map<String, String> map2) {
				
				int result;
				int date1 = Integer.parseInt(map1.get("regDate"));
				int date2 = Integer.parseInt(map2.get("regDate"));
				
				if(date1 > date2) result = -1;
				else if(date1==date2) result = 0;
				else result = 1;
				return result;
			}
			
		};
		
		String paramFromDate = param.get("fromDt").toString();
		String paramToDate = param.get("toDt").toString();
		String paramPath = param.get("path").toString();
		String paramType = param.get("type").toString();
		String paramAccCode = "";
		String paramUserId = "";
		String userPmsAccCode = "";
		
		filenameFilterUtil.setFromDt(paramFromDate);
		filenameFilterUtil.setToDt(paramToDate);
		
		if("1".equals(paramType)) {
			
			paramUserId = param.get("userId").toString();
			File tradeDir = new File(paramPath);
			if(tradeDir.isDirectory()) {
				
				fileFilterUtil.setNameList(paramUserId);
				File[] userTradeDirList = tradeDir.listFiles(fileFilterUtil);
				
				for(File userTradeDir : userTradeDirList) {
					
					String[] detailFileList = userTradeDir.list(filenameFilterUtil);
					String userId = userTradeDir.getName();
					
					for(String fileName : detailFileList) {
						HashMap<String, String> fileInfo = new HashMap<String, String>();
						fileInfo.put("fileName", fileName);
						fileInfo.put("filePath", userId + "/" + fileName);
						fileInfo.put("regDate", fileName.substring(fileName.lastIndexOf("_")+1, fileName.indexOf(".")));
						allFileList.add(fileInfo);
					}
					
				}
			}else if("2".equals(paramType)) {
				
				paramAccCode = param.get("accCode").toString();
				userPmsAccCode = param.get("userPmsAccCode").toString();
				
				File reportDir = new File(paramPath);
				
				if(reportDir.isDirectory()) {
					fileFilterUtil.setNameList(userPmsAccCode);
					fileFilterUtil.setSchAccCode(paramAccCode);
					
					File[] accCodeDirList = reportDir.listFiles(fileFilterUtil);
					
					for(File accCodeDir : accCodeDirList) {
						String accCodeName = accCodeDir.getName();
						filenameFilterUtil.setDirName(accCodeName+"_");
						String[] detailFileList = accCodeDir.list(filenameFilterUtil);
						
						for(String fileName : detailFileList) {
							HashMap<String, String> fileInfo = new HashMap<String, String>();
							fileInfo.put("fileName", fileName);
							fileInfo.put("regDate", fileName.substring(fileName.lastIndexOf("_")+1, fileName.indexOf(".")));
							fileInfo.put("filePath", accCodeName+"/"+fileName);
							fileInfo.put("fileAccCode", accCodeName);
							allFileList.add(fileInfo);
						}
					}
				}
				
			}
			
			//모든 파일 날짜순 정렬
			Collections.sort(allFileList, cprt);
			
			int totalSize = allFileList.size();
			
			for(int i=0; i<totalSize; i++) {
				allFileList.get(i).put("fileTotalCnt", Integer.toString(totalSize));
				allFileList.get(i).put("fileNo", Integer.toString(i+1));
			}
			
		}
		
		return allFileList;
	}
	
	public class FileFilterUtil implements FileFilter{

		private String type;
		private String nameList;
		private String schAccCode;
		
		public FileFilterUtil(String type) {
			this.type = type;
		}
		
		public String getNameList() {
			return nameList;
		}



		public void setNameList(String nameList) {
			this.nameList = nameList;
		}



		public String getSchAccCode() {
			return schAccCode;
		}



		public void setSchAccCode(String schAccCode) {
			this.schAccCode = schAccCode;
		}



		@Override
		public boolean accept(File file) {
			
			boolean result = false;
			
			//사용자 허용된 계좌 필터링
			if("report".equals(type)) {
				if(nameList.length()>0) {
					String[] checkNameList = nameList.split(",");
					for(String checkName : checkNameList) {
						if(file.getName().equals(checkName)) result = true;
					}
				}else {
					result = false;
				}
				
				if(!result) return result;
				
				//사용자가 검색한 계좌 코드
				if(this.schAccCode!= null && !"".equals(this.schAccCode)) {
					if(file.getName().equals(this.schAccCode)) result = true;
					else result = false;
				}
			}
			
			return result;
			//return 값이 true면 해당 파일 리스트에 추가
		}
		
	}
	
	public class FilenameFilterUtil implements FilenameFilter {

		private String type;
		private String fromDt;
		private String toDt;
		private String dirName;
		
		public FilenameFilterUtil(String type) {
			// TODO Auto-generated constructor stub
			this.type = type;
		}
		
		
		
		public String getFromDt() {
			return fromDt;
		}



		public void setFromDt(String fromDt) {
			this.fromDt = fromDt;
		}



		public String getToDt() {
			return toDt;
		}



		public void setToDt(String toDt) {
			this.toDt = toDt;
		}



		public String getDirName() {
			return dirName;
		}



		public void setDirName(String dirName) {
			this.dirName = dirName;
		}



		@Override
		public boolean accept(File file, String name) {
			boolean result = false;
			if("report".equals(type)) {
				int fromDate = Integer.parseInt(fromDt);
				int toDate = Integer.parseInt(toDt);
				int fileDate;
				
				if(!name.endsWith(".xlsx")) return false;
				
				if(dirName!=null) {
					String[] fileNameInfo = name.split(dirName);
					fileDate = Integer.parseInt(fileNameInfo[1].substring(0, fileNameInfo[1].indexOf(".")));
				}else {
					fileDate = Integer.parseInt(name.substring(name.lastIndexOf("_")+1, name.indexOf(".")));
				}
				
				if(fileDate >= fromDate && toDate >= fileDate) {
					result = true;
				}else {
					result = false;
				}
			}
			
			return result;
		}
		
	}
}
