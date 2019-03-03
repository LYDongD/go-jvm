package ch08;

public class ArrayDemo {

    public static void main(String[] args) {

        //数组构造指令
        int[] a1 = new int[10]; //newarray
        String[] a2 = new String[10]; //anewarray
        int[][] a3 = new int[10][10]; //multianewarray

        //数组长度指令
        int x = a1.length; //arraylength

        //数组的存储与加载指令
        a1[0] = 100; //iastore
        int y = a1[0]; //iaload
        a2[0] = "abc"; //aastore
        String s = a2[0]; //aaload

    }

}
