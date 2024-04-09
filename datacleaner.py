import csv

def extract_unique_words(input_csv, output_csv):
    unique_words = set()

    # Read input CSV file
    with open(input_csv, 'r', newline='') as csvfile:
        reader = csv.reader(csvfile)
        for row in reader:
            if row:  # Skip empty rows
                word = row[0].strip()  # Assuming the word is in the first column
                unique_words.add(word)

    # Write unique words to output CSV file
    with open(output_csv, 'w', newline='') as csvfile:
        writer = csv.writer(csvfile)
        for word in sorted(unique_words):  # Sort words alphabetically
            writer.writerow([word])

if __name__ == "__main__":
    input_csv = "dictionary.csv"
    output_csv = "database.csv"
    extract_unique_words(input_csv, output_csv)
    print("Unique words extracted and saved to", output_csv)
