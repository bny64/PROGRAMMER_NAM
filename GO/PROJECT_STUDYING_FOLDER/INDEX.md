1. project1 : 다양한 고 연습 폴더
2. project2 : revel을 사용한 고 웹프로그래밍
	*
	로그인 후 포스트 메인 화면으로 들어와서 로그아웃을 하게 되면 세션은 정상적으로 지워진다.
	다른 컨트롤러에서 c.CurrentUser를 출력해보면 nil이 찍히는데 
	setCurrentUser함수를 통과할 때 c.CurrentUser는 이전에 담아놓은 유저정보가 찍힌다.
	이해가 안간다...