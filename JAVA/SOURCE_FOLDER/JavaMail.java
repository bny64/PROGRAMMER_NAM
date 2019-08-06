public class javaMail(){
	public static void main(String[] args) throws Exception {
		Properties props = new Properties();
		props.put("mail.smtp.host", "host 주소");
		props.put("mail.smtp.port", "host port");
		
		Session session = Session.getDefaultInstance(props);
		session.setDebug(true);
		
		MimeMessage msg = new MimeMessage(session);
		msg.setFrom(new InternetAddress("보내는 사람 이메일", "이름"));
		msg.addRecipient(Message.RecipientType.TO, new InternetAddress("받는 사람 주소", "이름"));
		msg.setSubject("제목")
		msg.setContent("내용", "text/html; charset=utf-8");
		Transport.send(msg);		
	}
}