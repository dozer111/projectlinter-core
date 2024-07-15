package printer

// This is a package with a reusable code for projectlinter.Rule(FailedMessage)
// The main idea is
//
// There is 3 ways to show the failed message to user in projectlinter.Rule(FailedMessage)
// 1. Write own custom code in projectlinter.Rule(FailedMessage)
// 2. Use one of common reusable printers
// 3. Combine reusable printer and your own custom code
//
// As you see, I don`t do the Printer interface
// This is my architecture decision
// The main reason is - let the printer be maximum flexible
//
// Do what you want to do. I don`t care about the method(s) or structure. Just write the code in output correct
//
// The only limitations are
// 1. Each printer in separate file
// 2. Each printer MUST be tested
