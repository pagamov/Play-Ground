package main

// var n int = 999998727899999 * 21732382677641 // ~10^29
// var n int = 21732382677641 * 523347633027360537213687137 // ~10^40

// step to make in sieve process

// B = upper bound for Factor Base

// N = number to be factorised
// n = 160769357899975610828199539114109518167531134514190990785144666932076614717841
// n = 1250171497372227982026555999675170108947918951378367343470923483104158597216632066586300921566811265776465427395026458151240042366061271512107752586681699923914902061886213022544496783070727061083763996630816279869169194623169255711135422521925444135939014878277515299870536875962948267973899545621728547726545192382593936985574978881305949487523233148677106330650818223443955800622774189936635106363035784698216185461573761714766211607812695281252356674432444279

// done
// n = 97 * 13
// n = 104729 * 103591 # ~10^11

// var n int = 999998727899999 * 21732382677641 # ~10^29
// var n int = 21732382677641 * 523347633027360537213687137 # ~10^40
// n = 3331113965338635107 * 523347633027360537213687137 # ~10^45
// n = 618970019642690137449562111 * 523347633027360537213687137 # ~10^53
// n = 10848981839

// RSA 100
// n = 1522605027922533360535618378132637429718068114961380688657908494580122963258952897654000350692006139
// RSA 110
// n = 35794234179725868774991807832568455403003778024228226193532908190484670252364677411513516111204504060317568667

// do we need to search for factor base? default = True
var B_search bool = true

// do we need to save our factor base? default = False
var B_save bool = false

// if B_search == False, name file where we find factor base? default = primes.crypt
// file structure:
// data = [B, n, self.p, self.r]
var B_file string = "primes.crypt"

// do we need to find smooth numbers for given B? default = True
var Smooth_search bool = true

// use parallel search instead of serial
var Smooth_search_parallel bool = !Smooth_search

// do we need to save our smooth base? default = False
var Smooth_save bool = false

// if Smooth_search == False, name file where we find smooth numbers? default = smooth.crypt
// file structure:
// data = [B,n,smooth_numbers[x,q(x),vector]]
var Smooth_file string = "smooth.crypt"

var Parallel_max_processes int = 4

var Verbose_level int = 1

// """level of output detail"""

// do we need to form matrix? default = True
var Matrix_solve bool = true

// do we need to save our matrix? default = False
var matrix_save bool = false

// if Matrix_solve == False, name file where we find matrix? default = matrix.crypt
// file structure:
// B
// matrix for smooth numbers
var Matrix_file string = "matrix.crypt"
