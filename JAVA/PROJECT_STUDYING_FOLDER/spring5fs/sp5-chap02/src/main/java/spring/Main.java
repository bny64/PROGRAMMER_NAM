package spring;

public class Main {

	public static void main(String[] args) {
		// TODO Auto-generated method stub
		MemberDao memberDao = new MemberDao();
		MemberRegisterService regSvc = new MemberRegisterService(memberDao);
		ChangePasswordService pwdSvc= new ChangePasswordService();
		pwdSvc.setMemberDao(memberDao);
	}

}
