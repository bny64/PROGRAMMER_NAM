package chap02;

import org.springframework.context.annotation.AnnotationConfigApplicationContext;

public class Main {
	
	public static void main(String[] args) {
		
		/**
		 * AnnotationConfigApplicationContext : 자바 애노테이션을 이용한 클래스로부터 객체 설정 정보를 가져온다.
		 * GenericXmlApplicationContext : XML로부터 객체 설정 정보를 가져온다.
		 * GenericGroovyApplicationContext : 그루비 코드를 이용해 설정정보를 가져온다.
		 * */
		
		AnnotationConfigApplicationContext ctx = 
				new AnnotationConfigApplicationContext(AppContext.class);
		//자바 설정에서 정보를 읽어와 빈 객체를 생성하고 관리.
		//AnnotationConfigApplicationContext 객체를 생성할 때 앞서 작성한 AppContext클래스를 생성자 파라미터로 전달한다.
		Greeter g = ctx.getBean("greeter", Greeter.class);
		//getBean() AnnotationConfigApplicationContext가 자바 설정을 읽어와 생성한 빈 객체를 검색할 때 사용한다.
		//첫 번째 파라미터는 @Bean 애노테이션의 메서드 이름인 빈 객체의 이름.
		//두 번째는 빈 객체의 타입.
		//Greeter이 Greeter.class의 리턴타입이므로 타입은 Greeter.class
		String msg = g.greet("스프링");
		System.out.println(msg);
		ctx.close();
	}
	
}
