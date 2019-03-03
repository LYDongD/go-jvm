package ch08;

public class BubbleSort {

    public static void main(String[] args) {
        int[] arr = {22, 84, 57, 9,78, 1};
        bubbleSort(arr);
        printArray(arr);
    }

    private static void bubbleSort(int[] arr) {
        boolean swapped = true;
        int j = 0;
        int tmp;
        while (swapped) {
            swapped = false;
            j++;
            for (int i = 0; i < arr.length - j; i++) {
                if (arr[i] > arr[i+1]) {
                    tmp = arr[i];
                    arr[i+1] = arr[i];
                    arr[i+1] = tmp;
                }
                swapped = true;
            }
        }
    }

    private static void printArray(int[] arr) {
        for (int anArr : arr) {
            System.out.println(anArr);
        }
    }
}

