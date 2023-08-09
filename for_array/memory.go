package main

import "fmt"

/*
// C++
int main(int argc, char *argv[]) {
    std::string array[1];
    array[1] = "10000";

    std::cout << &array[0] << std::endl;
    std::cout << &array[1] << std::endl;
    std::cout << &array[2] << std::endl;
    std::cout << &array[3] << std::endl;
    std::cout << &array[4] << std::endl;
    std::cout << &array[5] << std::endl;

    return 0;
}
输出，间隔 24 bytes：
0x16b0ef150
0x16b0ef168
0x16b0ef180
0x16b0ef198
0x16b0ef1b0
*/

/*
输出，间隔 16 bytes：
0x1400007a010
0x1400007a020
0x1400007a030
0x1400007a040
0x1400007a050
0x1400007a060
0x1400007a070
0x1400007a080
0x1400007a090
*/
func main() {
	var array [10]string

	//array := make([]int, 10)
	array[1] = "kkk"
	//array[1] = 12

	fmt.Printf("%p\n", &array[0])
	fmt.Printf("%p\n", &array[1])
	fmt.Printf("%p\n", &array[2])
	fmt.Printf("%p\n", &array[3])
	fmt.Printf("%p\n", &array[4])
	fmt.Printf("%p\n", &array[5])
}
