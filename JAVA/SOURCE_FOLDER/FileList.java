package com.test;

import java.io.File;
import java.io.FileFilter;
import java.io.FilenameFilter;
import java.util.ArrayList;
import java.util.HashMap;
import java.util.List;
import java.util.Map;

public class FileList {
	public List<HashMap<String, String>> getFileList(Map param){
		List<HashMap<String, String>> allFileList = new ArrayList<HashMap<String, String>>();
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
