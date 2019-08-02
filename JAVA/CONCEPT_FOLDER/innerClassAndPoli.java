public class TestClass{
	public Class A {
		public void methodA(){
			System.out.println("methodA");
		}
	}
	public Class B extends A{
		public void methodB(){
			System.out.println("methodB");
		}
	}
	public Class C extends A{
		public void methodC(){
			System.out.println("methodC");
		}
	}
	
	public static void main(String[] args) {
		//innerClass 생성 방법.
		TestClass tc = new TestClass();
		TestClass.A a = tc.new A();
		TestClass.B b = tc.new B();
		TestClass.C c = tc.new C();
		a.methodA();
		b.methodA();
		b.methodB();
		c.methodA();
		c.methodC();
		TestClass.A aa = tc.new B();
		aa.methodA(); //A를 상속받는 B 클래스의 객체를 A타입으로 만들면 B클래스의 메서드는 사용할 수 없다.
	}
}