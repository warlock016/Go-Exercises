package csvparser

// ParseCSV parses a single line of CSV with quoted fields and escape sequences
func ParseCSV(line string) []string {
	// TODO(human): Implement the CSV parser
	//
	// Algorithm:
	// 1. Create result slice and buffer (strings.Builder)
	// 2. Track state: inQuotes (boolean)
	// 3. Use index-based loop: for i := 0; i < len(line); i++
	//
	// 4. Inside quotes (inQuotes == true):
	//    a. If char == '"':
	//       - Check if next char is also '"' (lookahead: i+1 < len(line) && line[i+1] == '"')
	//       - If yes: add single '"' to buffer, increment i to skip second quote
	//       - If no: set inQuotes = false (end of quoted field)
	//    b. Else: add char to buffer
	//
	// 5. Outside quotes (inQuotes == false):
	//    a. If char == ',': flush buffer to result, reset buffer
	//    b. If char == '"': set inQuotes = true (start of quoted field)
	//    c. Else: add char to buffer
	//
	// 6. After loop: flush final field to result
	//
	// Hint: Use buffer.WriteByte(char) for single bytes
	// Hint: Use buffer.WriteRune('"') for the escaped quote character
	// Hint: Remember to check bounds before lookahead: i+1 < len(line)
	// Hint: You'll need to import "strings" package for strings.Builder

	return nil // TODO(human): Replace with your implementation
}
