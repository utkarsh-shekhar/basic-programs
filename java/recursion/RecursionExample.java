public class RecursionExample {

	public static void main(String[] args) {
		System.out.println(addAllNumbersPrevious(10));
	}
	
	private static int addAllNumbersPrevious(int x) {
		if(x <= 1) {
			return 1;
		}
		else {
			return x + addAllNumbersPrevious(x - 1);
		}
	}

}
