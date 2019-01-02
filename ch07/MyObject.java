package gojvm.ch07;

public class MyObject{
    public static int staticVar;
    public int instanceVar;
    public static void main(String[] args) {
        int x = 32768; //ldc
        MyObject myObject = new MyObject(); //new
        MyObject.staticVar = x; // putstatic
        x = MyObject.staticVar; //getstatic
        myObject.instanceVar = x; //putfield
        x = myObject.instanceVar; //getfield
        Object obj = myObject;
        if (obj instanceof MyObject) { //instanceof
            myObject = (MyObject)obj; //checkcast
            System.out.println(myObject.instanceVar);
        }
    }
}