package main

import (
	"fmt"
)

func byte2string(conversionCandidate interface{}) {
	fmt.Println(string(conversionCandidate.([]byte)))
}

func main() {
	var conversionCandidate []byte
	conversionCandidate = []byte{123,}//[123 34 101 118 101 110 116 86 101 114 115 105 111 110 34 58 34 49 46 48 53 34 44 34 117 115 101 114 73 100 101 110 116 105 116 121 34 58 123 34 116 121 112 101 34 58 34 65 115 115 117 109 101 100 82 111 108 101 34 44 34 112 114 105 110 99 105 112 97 108 73 100 34 58 34 65 82 79 65 74 84 68 65 85 80 55 69 71 84 67 50 72 84 68 82 52 58 100 111 99 115 95 112 104 111 116 111 105 100 95 102 97 99 105 97 108 95 114 101 107 111 34 44 34 97 114 110 34 58 34 97 114 110 58 97 119 115 58 115 116 115 58 58 51 51 48 57 51 50 54 52 55 56 51 54 58 97 115 115 117 109 101 100 45 114 111 108 101 47 100 111 99 115 95 112 104 111 116 111 105 100 95 102 97 99 105 97 108 95 114 101 107 111 95 108 97 109 98 100 97 95 114 111 108 101 47 100 111 99 115 95 112 104 111 116 111 105 100 95 102 97 99 105 97 108 95 114 101 107 111 34 44 34 97 99 99 111 117 110 116 73 100 34 58 34 51 51 48 57 51 50 54 52 55 56 51 54 34 44 34 97 99 99 101 115 115 75 101 121 73 100 34 58 34 65 83 73 65 85 50 68 73 55 79 54 79 73 55 69 72 89 67 71 74 34 44 34 115 101 115 115 105 111 110 67 111 110 116 101 120 116 34 58 123 34 97 116 116 114 105 98 117 116 101 115 34 58 123 34 109 102 97 65 117 116 104 101 110 116 105 99 97 116 101 100 34 58 34 102 97 108 115 101 34 44 34 99 114 101 97 116 105 111 110 68 97 116 101 34 58 34 50 48 49 56 45 48 56 45 48 51 84 49 52 58 51 51 58 53 55 90 34 125 44 34 115 101 115 115 105 111 110 73 115 115 117 101 114 34 58 123 34 116 121 112 101 34 58 34 82 111 108 101 34 44 34 112 114 105 110 99 105 112 97 108 73 100 34 58 34 65 82 79 65 74 84 68 65 85 80 55 69 71 84 67 50 72 84 68 82 52 34 44 34 97 114 110 34 58 34 97 114 110 58 97 119 115 58 105 97 109 58 58 51 51 48 57 51 50 54 52 55 56 51 54 58 114 111 108 101 47 100 111 99 115 95 112 104 111 116 111 105 100 95 102 97 99 105 97 108 95 114 101 107 111 95 108 97 109 98 100 97 95 114 111 108 101 34 44 34 97 99 99 111 117 110 116 73 100 34 58 34 51 51 48 57 51 50 54 52 55 56 51 54 34 44 34 117 115 101 114 78 97 109 101 34 58 34 100 111 99 115 95 112 104 111 116 111 105 100 95 102 97 99 105 97 108 95 114 101 107 111 95 108 97 109 98 100 97 95 114 111 108 101 34 125 125 125 44 34 101 118 101 110 116 84 105 109 101 34 58 34 50 48 49 56 45 48 56 45 48 51 84 49 53 58 51 48 58 50 53 90 34 44 34 101 118 101 110 116 83 111 117 114 99 101 34 58 34 101 99 50 46 97 109 97 122 111 110 97 119 115 46 99 111 109 34 44 34 101 118 101 110 116 78 97 109 101 34 58 34 67 114 101 97 116 101 78 101 116 119 111 114 107 73 110 116 101 114 102 97 99 101 34 44 34 97 119 115 82 101 103 105 111 110 34 58 34 117 115 45 101 97 115 116 45 49 34 44 34 115 111 117 114 99 101 73 80 65 100 100 114 101 115 115 34 58 34 53 52 46 50 50 49 46 49 49 46 51 57 34 44 34 117 115 101 114 65 103 101 110 116 34 58 34 97 119 115 45 105 110 116 101 114 110 97 108 47 51 32 97 119 115 45 115 100 107 45 106 97 118 97 47 49 46 49 49 46 51 53 52 32 76 105 110 117 120 47 52 46 49 52 46 53 49 45 54 48 46 51 56 46 97 109 122 110 49 46 120 56 54 95 54 52 32 74 97 118 97 95 72 111 116 83 112 111 116 40 84 77 41 95 54 52 45 66 105 116 95 83 101 114 118 101 114 95 86 77 47 50 53 46 49 55 50 45 98 51 49 32 106 97 118 97 47 49 46 56 46 48 95 49 55 50 34 44 34 114 101 113 117 101 115 116 80 97 114 97 109 101 116 101 114 115 34 58 123 34 115 117 98 110 101 116 73 100 34 58 34 115 117 98 110 101 116 45 55 101 56 48 102 57 50 55 34 44 34 100 101 115 99 114 105 112 116 105 111 110 34 58 34 65 87 83 32 76 97 109 98 100 97 32 86 80 67 32 69 78 73 58 32 100 97 99 100 55 100 57 50 45 53 57 55 53 45 52 49 55 97 45 97 49 57 97 45 48 57 102 55 52 99 102 50 51 56 57 97 34 44 34 103 114 111 117 112 83 101 116 34 58 123 34 105 116 101 109 115 34 58 91 123 34 103 114 111 117 112 73 100 34 58 34 115 103 45 101 100 50 97 57 57 57 50 34 125 93 125 44 34 112 114 105 118 97 116 101 73 112 65 100 100 114 101 115 115 101 115 83 101 116 34 58 123 125 125 44 34 114 101 115 112 111 110 115 101 69 108 101 109 101 110 116 115 34 58 123 34 114 101 113 117 101 115 116 73 100 34 58 34 48 57 50 99 98 51 51 99 45 52 53 49 101 45 52 99 57 55 45 57 52 99 100 45 102 55 100 57 54 51 99 55 52 48 54 100 34 44 34 110 101 116 119 111 114 107 73 110 116 101 114 102 97 99 101 34 58 123 34 110 101 116 119 111 114 107 73 110 116 101 114 102 97 99 101 73 100 34 58 34 101 110 105 45 48 101 102 98 101 52 51 101 34 44 34 115 117 98 110 101 116 73 100 34 58 34 115 117 98 110 101 116 45 55 101 56 48 102 57 50 55 34 44 34 118 112 99 73 100 34 58 34 118 112 99 45 50 51 51 51 50 53 52 54 34 44 34 97 118 97 105 108 97 98 105 108 105 116 121 90 111 110 101 34 58 34 117 115 45 101 97 115 116 45 49 99 34 44 34 100 101 115 99 114 105 112 116 105 111 110 34 58 34 65 87 83 32 76 97 109 98 100 97 32 86 80 67 32 69 78 73 58 32 100 97 99 100 55 100 57 50 45 53 57 55 53 45 52 49 55 97 45 97 49 57 97 45 48 57 102 55 52 99 102 50 51 56 57 97 34 44 34 111 119 110 101 114 73 100 34 58 34 51 51 48 57 51 50 54 52 55 56 51 54 34 44 34 114 101 113 117 101 115 116 101 114 73 100 34 58 34 65 82 79 65 74 84 68 65 85 80 55 69 71 84 67 50 72 84 68 82 52 58 100 111 99 115 95 112 104 111 116 111 105 100 95 102 97 99 105 97 108 95 114 101 107 111 34 44 34 114 101 113 117 101 115 116 101 114 77 97 110 97 103 101 100 34 58 102 97 108 115 101 44 34 115 116 97 116 117 115 34 58 34 112 101 110 100 105 110 103 34 44 34 109 97 99 65 100 100 114 101 115 115 34 58 34 48 101 58 55 98 58 48 54 58 48 97 58 99 50 58 101 52 34 44 34 112 114 105 118 97 116 101 73 112 65 100 100 114 101 115 115 34 58 34 49 48 48 46 49 48 53 46 49 51 53 46 53 51 34 44 34 112 114 105 118 97 116 101 68 110 115 78 97 109 101 34 58 34 105 112 45 49 48 48 45 49 48 53 45 49 51 53 45 53 51 46 101 99 50 46 105 110 116 101 114 110 97 108 34 44 34 115 111 117 114 99 101 68 101 115 116 67 104 101 99 107 34 58 116 114 117 101 44 34 105 110 116 101 114 102 97 99 101 84 121 112 101 34 58 34 105 110 116 101 114 102 97 99 101 34 44 34 103 114 111 117 112 83 101 116 34 58 123 34 105 116 101 109 115 34 58 91 123 34 103 114 111 117 112 73 100 34 58 34 115 103 45 101 100 50 97 57 57 57 50 34 44 34 103 114 111 117 112 78 97 109 101 34 58 34 104 122 115 118 99 45 108 97 109 98 100 97 45 115 103 34 125 93 125 44 34 112 114 105 118 97 116 101 73 112 65 100 100 114 101 115 115 101 115 83 101 116 34 58 123 34 105 116 101 109 34 58 91 123 34 112 114 105 118 97 116 101 73 112 65 100 100 114 101 115 115 34 58 34 49 48 48 46 49 48 53 46 49 51 53 46 53 51 34 44 34 112 114 105 118 97 116 101 68 110 115 78 97 109 101 34 58 34 105 112 45 49 48 48 45 49 48 53 45 49 51 53 45 53 51 46 101 99 50 46 105 110 116 101 114 110 97 108 34 44 34 112 114 105 109 97 114 121 34 58 116 114 117 101 125 93 125 44 34 105 112 118 54 65 100 100 114 101 115 115 101 115 83 101 116 34 58 123 125 44 34 116 97 103 83 101 116 34 58 123 125 125 125 44 34 114 101 113 117 101 115 116 73 68 34 58 34 48 57 50 99 98 51 51 99 45 52 53 49 101 45 52 99 57 55 45 57 52 99 100 45 102 55 100 57 54 51 99 55 52 48 54 100 34 44 34 101 118 101 110 116 73 68 34 58 34 99 54 97 97 98 98 49 51 45 100 49 100 54 45 52 52 54 101 45 98 56 98 51 45 50 54 55 55 100 53 56 102 99 101 100 53 34 44 34 101 118 101 110 116 84 121 112 101 34 58 34 65 119 115 65 112 105 67 97 108 108 34 125] 
	byte2string(conversionCandidate)
}
