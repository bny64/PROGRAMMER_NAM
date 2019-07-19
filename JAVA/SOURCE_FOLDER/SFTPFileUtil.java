public class SFTPFileUtil {
    private Session session = null;
    private Channel channel = null;
    private ChannelSftp channelSftp = null;

    public void ttsInit(String host, String userName, String password, int port) throws Exception {
        JSch jsch = new JSch();
        String[] hosts = host.split("\\|");

        try{
            session = jsch.getSession(userName, hosts[0], port);
            session.setPassword(password);

            java.util.Properties config = new java.util.Properties();
            config.put("StrictHostKeyChecking", "no");

            session.setConfig(config);
            session.setTimeout(5000);
            session.connect();

            channel = session.openChannel("sftp");
            channel.connect();

        }catch(SHchException e1){
            e1.printStackTrace();
            session = null;
            channel = null;

            try {
                
                session = jsch.getSession(userName, hosts[1], port);
                session.setPassword(password);

                java.util.Properties config = new java.util.Properties();
                config.put("StrictHostKeyChecking", "no");

                session.setConfig(config);
                session.setTimeout(5000);
                session.connect();

                channel = session.openChannel("sftp");
                channel.connect();
            }catch(JSchException e2){
                e2.printStackTrace();
                throw new Exception(e2);
            }
        }

        channelSftp = (ChannelSftp) channel;
    }

    public void fileUpload(String dir, String[] dateDir, File file){
        FileInputStream in = null;
        SftpATTRS attrs;

        for(int i=0; i<dateDir.length; i++){
            attrs = null;
            dir += "/" + dateDir[i];

            try{
                attrs = channelSftp.stat(dir);

            }catch(SftpException e){
                System.out.println("NOT FOUND");
            }

            if(attrs = null){
                try {
                    channelSftp.mkdir(dir);
                }catch(SftpException e){
                    
                }
            }
        }

        try {
            in = new FileInputStream(file);
            channelSftp.cd(dir);
            channelSftp.put(in, file.getName());
        }catch(FileNotFoundException e){

        }catch(SftpException e){

        }finally {
            try {
                in.close();
            }catch(IOException e){

            }
        }
    }

    public void deleteFile(String fileDir){
        SftpATTRS attrs = null;

        try {
            attrs = channelSftp.stat(fileDir);

            if(attrs !=null){
                channelSftp.rm(fileDir);
            }
        }catch(SftpException e) {

        }
    }

    public String checkFileExist(String dir, String savedDir, String publicDate, String fileName){
        SftpATTRS attrs = null;
        String fileYear = publicDate.substring(0,4);
        String fileMonth = publicDate.substring(4,6);
        String fileDay = publicDate.substring(6,8);
        String[] fileDate = {fileYear, fileMonth, fileDay};
        
        InputStream in = null;
		boolean result = false;
		
		for(int i=0; i<fileDate.length; i++){
            attrs = null;
			//logger.info("@@@@@ dir1 : {}", dir);
            dir += "/" + fileDate[i];
            //logger.info("@@@@@ dir2 : {}", dir);
			try {
				attrs = channelSftp.stat(dir);
			} catch (SftpException e) {
				logger.debug("cannot find directory!");
				return "";
			}
			
		}
		
		try {
			//logger.info("@@@@@ dir : {}", dir);
			channelSftp.cd(dir);
			in = channelSftp.get(fileName);
		} catch (SftpException e) {
			logger.debug("cannot find file!");
			return "";
		}
		
		try {
			
			int i;
			
			while((i = in.read()) != -1){
				result = true;
			}			
			
		} catch (IOException e) {
			// TODO Auto-generated catch block
			e.printStackTrace();
		}finally{
			try {
				in.close();
			} catch (IOException e) {
				// TODO Auto-generated catch block
				e.printStackTrace();
			}
		}
		
		if(result){
			return savedDir + "/" + fileYear + "/" + fileMonth + "/" + fileDay + "/" + fileName;
		}else{
			return "";
		}
    }

    public void disConnection(){
        channelSftp.quit();
    }
}