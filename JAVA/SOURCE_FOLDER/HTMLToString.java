public class StringUtil {

	public static void main(String[] args) throws Exception {
		String path = "D:/...path";
		String result = getHTML(path);
		System.out.Println(result);
	}
	
	public static String getHTML(String HtmlToRead){
		
		StringBuilder contentBuilder = new StringBuilder();
		try {
			BufferedReader in = new BufferedReader(new InputStreamReader(new FileInputStream(HtmlToRead, "UTF-8")));
			String str;
			while((str = in.readLine()) != null){
				contentBuilder.append(str);
			}
			in.close();
		}catch(IOException e){
			e.printStackTrace();
		}
		return contentBuilder.toString();
	}

}