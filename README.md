# Go
To run a .go file use below syntax

go run file_name.go

Ex.

    go run HelloWorld.go

To build and executable file for Go file use below syntax

go build file_name.go

Ex.

    go build HelloWorld.go

## Data Types

-   Numeric - 
&emsp;&emsp;int
&emsp;&emsp;float

-   Boolean - 
&emsp;&emsp;true
&emsp;&emsp;false

-   String

-   Derived - 
&emsp;&emsp;Pointer
&emsp;&emsp;Array
&emsp;&emsp;Structure
&emsp;&emsp;Map
&emsp;&emsp;Interface

## Operators

<table>
    <tr>
        <th style="text-align:center">Arithmetic</th>
        <th style="text-align:center">Relational</th>
        <th style="text-align:center">Logical</th>
    </tr>
    <tr>
        <td>+ Addition</td>
        <td >> Grater than</td>
        <td>&& Logical and</td>
    </tr>
    <tr>
        <td>- Subtraction</td>
        <td>< Lesser than</td>
        <td>|| Logical or</td>
    </tr>
    <tr>
        <td>* Multiplication</td>
        <td>>= Greater than or Equal to</td>
        <td>! Negation</td>
    </tr>
    <tr>
        <td>/ Division</td>
        <td>== Equivalence</td>
        <td></td>
    </tr>
    <tr>
        <td>% Modulus</td>
        <td>!= Not equals</td>
        <td></td>
    </tr>
</table>

## Format Printing

```
        fmt.Printf("%.3f \n", pi) //Returns three decimal or floating values
	fmt.Printf("%T \n", name) //Returns Data type
	fmt.Printf("%t \n", win)  //Return Boolean
	fmt.Printf("%d \n", x) //Returns Integer
	fmt.Printf("%b \n", 3) //Returns Binaries
	fmt.Printf("%c \n", 3) //Returns ASCII
	fmt.Printf("%x \n", 15) //Returns Hex Code
	fmt.Printf("%e \n", pi) //Retuens Scientific notations
```