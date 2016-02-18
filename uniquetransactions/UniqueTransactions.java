import java.util.List;
import java.util.Arrays;
import java.util.Collections;
import java.util.stream.Collectors;

public class UniqueTransactions {
	public static void main(String[] args) {
		List<Transaction> transactions = Arrays.asList(
			new Transaction(100), 
			new Transaction(-100), 
			new Transaction(200)
		);

		transactions.stream()
				.filter(i -> Collections.frequency(transactions, i) == 1)
                .collect(Collectors.toSet())
                .forEach(System.out::println);
	}
}

class Transaction {
	private double value;

	public Transaction(double value) {
		this.value = value;
	}

	public double getValue() {
		return value;
	}

	public String toString() {
		return "Transaction(value=" + this.value + ")";
	}

	@Override
	public boolean equals(Object o) {
		Transaction other = (Transaction)o;

		if (this.value == Math.abs(other.getValue())) return true;

		return false;
	}
}