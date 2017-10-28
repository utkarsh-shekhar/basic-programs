public class Student {
	//MARK: Instance Variables
	private String studentMajor;
	private String firstName;
	private String lastName;
	private int studentCredits;
	private double studentGPA;
	
	//MARK: Constructors
	public Student(String fName, String lName) {
		this.firstName = fName;
		this.lastName = lName;
		this.studentMajor = "General Studies";
		this.studentGPA = 0;
		this.studentCredits = 0;
	}
	
	public Student(String major, int credits, double points, String fName, String lName) {
		this.studentMajor = major;
		this.studentCredits = credits;
		this.studentGPA = points;
		this.firstName = fName;
		this.lastName = lName;
	}

	//MARK: Getters and Setters
	public String getStudentMajor() {
		return studentMajor;
	}

	public String getFullName() {
		return firstName + " " + lastName;
	}

	public int getStudentCredits() {
		return studentCredits;
	}

	public double getStudentGPA() {
		return studentGPA;
	}
	
	public void ChangeMajor(String newMajor) {
		this.studentMajor = newMajor;
	}
}
