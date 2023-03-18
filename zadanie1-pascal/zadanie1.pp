program zadanie1;

uses sysutils;
type integerArrayType = array of integer;
type booleanArrayType = array of boolean;
var
    randomNumbersArray: integerArrayType;
    testResultsArray: booleanArrayType;
    ii: integer;

(* procedure returning count numbers in range 
var result is passed by reference *)
procedure randomNumbers(startRange, endRange, count: integer; var result: integerArrayType);
var valueCopy, i: integer;
begin
    if count <= 0 then
    begin
        writeln('Array size cannot be smaller than 1.');
        exit;
    end;
    if startRange > endRange then
    begin
        valueCopy := startRange;
        startRange := endRange;
        endRange := valueCopy;
    end;
    
    setlength(result, count);
    for i:= 0 to count-1 do
        result[i] := random(endRange - startRange + 1) + startRange;
end;

// procedure sorting int array
procedure sortNumbers(var intArray: integerArrayType);
var valueCopy, i, j: integer;
begin
    if high(intArray) <= low(intArray) then
    begin
        writeln('Array size to sort cannot be smaller than 2.');
        exit;
    end;
    
    for i:= low(intArray) to high(intArray)-1 do
        for j:= low(intArray) to high(intArray)-1-i do
            if intArray[j] > intArray[j+1] then
            begin
                valueCopy := intArray[j];
                intArray[j] := intArray[j+1];
                intArray[j+1] := valueCopy;
            end;
end;

function testWrongArraySize(): boolean;
var
    testArr: integerArrayType;
    result: boolean = false;
begin
    writeln('===== TEST: testWrongArraySize =====');
    randomNumbers(1, -1, -2, testArr);
    sortNumbers(testArr);
    if high(testArr) <= 0 then
        result := true;
        
    writeln('Expected length: <= 0');
    writeln('Actual length: ', high(testArr));
    writeln('Passed: ', result);
    testWrongArraySize := result;
end;

function testNegativeNumbers(): boolean;
var 
    testArr: integerArrayType;
    result: boolean = true;
    i: integer;
    strResult: ansistring = '';
const length: integer = 10;
begin
    writeln('===== TEST: testNegativeNumbers =====');
    randomNumbers(-55, -1, length, testArr);
    for i := low(testArr) to high(testArr) do
        begin
        strResult := strResult + inttostr(testArr[i]) + ', ';
        if testArr[i] >= 0 then
            result := false;
        end;
    
    writeln('Expected only values lower than 0');
    writeln('Found value >= 0: ', not result);
    writeln('Values: ', strResult);
    writeln('Passed: ', result);
    testNegativeNumbers := result;
end;

function testSorting(): boolean;
var
    testArr: integerArrayType = (-1, 6, -222, 0, 66, -2, 1);
    testArrExpected: integerArrayType = (-222, -2, -1, 0, 1, 6, 66);
    result: boolean = true;
    i: integer;
    strTestArr: ansistring = '';
    strTestArrExpected: ansistring = '';
    strResult: ansistring = '';
begin
    writeln('===== TEST: testSorting =====');
    for i := low(testArr) to high(testArr) do
    begin
        strTestArr := strTestArr + inttostr(testArr[i]) + ', ';
        strTestArrExpected := strTestArrExpected + inttostr(testArrExpected[i]) + ', ';
    end;
    
    sortNumbers(testArr);
    for i := low(testArr) to high(testArr) do
        begin
        strResult := strResult + inttostr(testArr[i]) + ', ';
        if testArr[i] <> testArrExpected[i] then
            result := false;
        end;
        
    writeln('Initial values: ');
    writeln(strTestArr);
    writeln('Expected values: ');
    writeln(strTestArrExpected);
    writeln('Result values: ');
    writeln(strResult);
        
    writeln('Passed: ', result);
    testSorting := result;
end;

function testCountGeneratedNumbers(): boolean;
var
    result: boolean = false;
    testArr: integerArrayType;
    resultLength, i: integer;
    strResult: ansistring = '';
const
    length: integer = 33;
begin
    writeln('===== TEST: testCountGeneratedNumbers =====');
    randomNumbers(-9, 9, length, testArr);
    writeln('Expected length of array: ', inttostr(length));
    resultLength := high(testArr) - low(testArr) + 1;
    writeln('Actual length of array: ', inttostr(resultLength));
    if resultLength = length then
        result := true;
    
    for i := low(testArr) to high(testArr) do
        strResult := strResult + inttostr(testArr[i]) + ', ';
    
    writeln('Array: ', strResult);
    writeln('Passed: ', result);
    testCountGeneratedNumbers := result;
end;

function testInvertedStartEnd(): boolean;
var
    result: boolean = true;
    testArr: integerArrayType;
    i: integer;
    strResult: ansistring = '';
const
    startRange: integer = 1;
    endRange: integer = 3333;
    length: integer = 33;
begin
    randomNumbers(endRange, startRange, length, testArr);
    for i := low(testArr) to high(testArr) do
    begin
        strResult := strResult + inttostr(testArr[i]) + ', ';
        if (testArr[i] > endRange) or (testArr[i] < startRange) then
            result := false;
    end;
    if low(testArr)+high(testArr)+1 <> length then
        result := false;
    
    writeln('Array: ', strResult);
    writeln('Values should be in range <', inttostr(startRange), ', ', inttostr(endRange), '>');
    writeln('Passed: ', result);
    testInvertedStartEnd := result;
end;

begin
    // initializes random numbers - without it they weren't looking random
    randomize();
    randomNumbers(1, -1, 0, randomNumbersArray);
    sortNumbers(randomNumbersArray);
    
    // tests
    setlength(testResultsArray, 5);
    testResultsArray[1] := testWrongArraySize();
    testResultsArray[2] := testNegativeNumbers();
    testResultsArray[3] := testSorting();
    testResultsArray[4] := testCountGeneratedNumbers();
    testResultsArray[5] := testInvertedStartEnd();
    
    writeln('Summary:');
    for ii := 1 to 5 do
        writeln('Test ', inttostr(ii), ' result: ', testResultsArray[ii]);
end.
