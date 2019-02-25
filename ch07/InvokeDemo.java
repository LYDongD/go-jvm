package gojvm.ch07;

public class InvokeDemo implements Runnable{

    public static void staticMethod(){}
    private void instanceMethod() {}
    @Override
    public void run() {
    }

    public void test() {
        InvokeDemo.staticMethod(); //invokestatic
        InvokeDemo invokeDemo = new InvokeDemo(); //invokespecial
        invokeDemo.instanceMethod(); //invokevirtual
        super.equals(null);//invokespecial
        this.run();//invokevirtual
        ((Runnable)invokeDemo).run();//invokeinterface
    }

    public static void main(String[] args) {
        new InvokeDemo().test();
    }
}