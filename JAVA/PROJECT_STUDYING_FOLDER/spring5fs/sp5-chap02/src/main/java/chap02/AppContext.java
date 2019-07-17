package chap02;

import org.springframework.context.annotation.Bean;
import org.springframework.context.annotation.Configuration;

@Configuration //스프링 설정 클래스
public class AppContext {
	
	@Bean //해당 메서드가 생성한 객체를 스프링이 관리하는 빈 객체로 등록.
	//애노테이션을 붙인 메서드의 이름은 빈 객체를 구분할 때 사용함.
	//@Bean애노테이션을 붙인 메서드는 객체를 생성하고 알맞게 초기화해야함.
	public Greeter greeter() {
		Greeter g = new Greeter();
		g.setFormat("%s, 안녕하세요");
		return g;
	}
}
