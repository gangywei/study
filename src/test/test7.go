package test

//两个水瓶倒水问题，求最大公约数，能被最大公约数整除就代表可以
func CanMeasureWater(x int, y int, z int) bool {
	if x+y < z {
		return false
	}
	if x == 0 || y == 0 {
		return z == 0 || x+y == z
	}
	gcdNumber := Gcd(x, y)
	if z%gcdNumber == 0 {
		return true
	}
	return false
}

//https://www.cnblogs.com/fusiwei/p/11301436.html
func Gcd(x, y int) int {
	if x > y {
		y, x = x, y
	}
	remain := y % x
	if remain == 0 {
		return x
	}
	return Gcd(x, remain)
}
