// package hello; // Error: Could not find or load main class Main -- 当直接在本目录运行时
package hello; // 添加包名的话，就不能直接在本目录运行；必须在上层目录运行

public class Main { // MainApp.java:1: error: class Main is public, should be declared in a file named Main.java
    public static void main(String[] args) {  
        System.out.println("hello world");
    }  
} 