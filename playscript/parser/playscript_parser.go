// Code generated from PlayScript.g4 by ANTLR 4.8. DO NOT EDIT.

package parser // PlayScript

import (
	"fmt"
	"reflect"
	"strconv"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = reflect.Copy
var _ = strconv.Itoa

var parserATN = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 115, 589,
	4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7,
	4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12, 4, 13,
	9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 4, 18, 9,
	18, 4, 19, 9, 19, 4, 20, 9, 20, 4, 21, 9, 21, 4, 22, 9, 22, 4, 23, 9, 23,
	4, 24, 9, 24, 4, 25, 9, 25, 4, 26, 9, 26, 4, 27, 9, 27, 4, 28, 9, 28, 4,
	29, 9, 29, 4, 30, 9, 30, 4, 31, 9, 31, 4, 32, 9, 32, 4, 33, 9, 33, 4, 34,
	9, 34, 4, 35, 9, 35, 4, 36, 9, 36, 4, 37, 9, 37, 4, 38, 9, 38, 4, 39, 9,
	39, 4, 40, 9, 40, 4, 41, 9, 41, 4, 42, 9, 42, 4, 43, 9, 43, 4, 44, 9, 44,
	4, 45, 9, 45, 4, 46, 9, 46, 4, 47, 9, 47, 4, 48, 9, 48, 4, 49, 9, 49, 3,
	2, 3, 2, 3, 2, 3, 2, 5, 2, 103, 10, 2, 3, 2, 3, 2, 5, 2, 107, 10, 2, 3,
	2, 3, 2, 3, 3, 3, 3, 7, 3, 113, 10, 3, 12, 3, 14, 3, 116, 11, 3, 3, 3,
	3, 3, 3, 4, 3, 4, 5, 4, 122, 10, 4, 3, 5, 3, 5, 3, 5, 5, 5, 127, 10, 5,
	3, 6, 5, 6, 130, 10, 6, 3, 6, 3, 6, 3, 6, 3, 6, 7, 6, 136, 10, 6, 12, 6,
	14, 6, 139, 11, 6, 3, 6, 3, 6, 5, 6, 143, 10, 6, 3, 6, 3, 6, 3, 7, 3, 7,
	5, 7, 149, 10, 7, 3, 8, 3, 8, 5, 8, 153, 10, 8, 3, 9, 3, 9, 3, 9, 7, 9,
	158, 10, 9, 12, 9, 14, 9, 161, 11, 9, 3, 10, 3, 10, 5, 10, 165, 10, 10,
	3, 10, 3, 10, 3, 11, 3, 11, 3, 11, 7, 11, 172, 10, 11, 12, 11, 14, 11,
	175, 11, 11, 3, 11, 3, 11, 5, 11, 179, 10, 11, 3, 11, 5, 11, 182, 10, 11,
	3, 12, 7, 12, 185, 10, 12, 12, 12, 14, 12, 188, 11, 12, 3, 12, 3, 12, 3,
	12, 3, 13, 7, 13, 194, 10, 13, 12, 13, 14, 13, 197, 11, 13, 3, 13, 3, 13,
	3, 13, 3, 13, 3, 14, 3, 14, 3, 15, 3, 15, 3, 15, 7, 15, 208, 10, 15, 12,
	15, 14, 15, 211, 11, 15, 3, 16, 3, 16, 3, 16, 3, 17, 3, 17, 3, 17, 3, 17,
	5, 17, 220, 10, 17, 3, 17, 3, 17, 3, 18, 3, 18, 3, 18, 3, 18, 7, 18, 228,
	10, 18, 12, 18, 14, 18, 231, 11, 18, 3, 19, 3, 19, 3, 19, 5, 19, 236, 10,
	19, 3, 20, 3, 20, 3, 20, 7, 20, 241, 10, 20, 12, 20, 14, 20, 244, 11, 20,
	3, 21, 3, 21, 5, 21, 248, 10, 21, 3, 22, 3, 22, 3, 22, 3, 22, 7, 22, 254,
	10, 22, 12, 22, 14, 22, 257, 11, 22, 3, 22, 5, 22, 260, 10, 22, 5, 22,
	262, 10, 22, 3, 22, 3, 22, 3, 23, 3, 23, 3, 23, 7, 23, 269, 10, 23, 12,
	23, 14, 23, 272, 11, 23, 3, 24, 3, 24, 3, 24, 3, 24, 5, 24, 278, 10, 24,
	5, 24, 280, 10, 24, 3, 25, 3, 25, 3, 25, 3, 25, 3, 25, 3, 25, 5, 25, 288,
	10, 25, 3, 26, 3, 26, 3, 27, 3, 27, 3, 28, 3, 28, 3, 29, 3, 29, 3, 29,
	3, 29, 3, 30, 7, 30, 301, 10, 30, 12, 30, 14, 30, 304, 11, 30, 3, 31, 3,
	31, 3, 31, 3, 31, 3, 31, 3, 31, 5, 31, 312, 10, 31, 3, 32, 3, 32, 3, 32,
	3, 32, 3, 32, 3, 32, 5, 32, 320, 10, 32, 3, 32, 3, 32, 3, 32, 3, 32, 3,
	32, 3, 32, 3, 32, 3, 32, 3, 32, 3, 32, 3, 32, 3, 32, 3, 32, 3, 32, 3, 32,
	3, 32, 3, 32, 3, 32, 3, 32, 3, 32, 7, 32, 342, 10, 32, 12, 32, 14, 32,
	345, 11, 32, 3, 32, 7, 32, 348, 10, 32, 12, 32, 14, 32, 351, 11, 32, 3,
	32, 3, 32, 3, 32, 3, 32, 5, 32, 357, 10, 32, 3, 32, 3, 32, 3, 32, 5, 32,
	362, 10, 32, 3, 32, 3, 32, 3, 32, 5, 32, 367, 10, 32, 3, 32, 3, 32, 3,
	32, 3, 32, 3, 32, 3, 32, 3, 32, 3, 32, 5, 32, 377, 10, 32, 3, 33, 6, 33,
	380, 10, 33, 13, 33, 14, 33, 381, 3, 33, 6, 33, 385, 10, 33, 13, 33, 14,
	33, 386, 3, 34, 3, 34, 3, 34, 5, 34, 392, 10, 34, 3, 34, 3, 34, 3, 34,
	5, 34, 397, 10, 34, 3, 35, 3, 35, 5, 35, 401, 10, 35, 3, 35, 3, 35, 5,
	35, 405, 10, 35, 3, 35, 3, 35, 5, 35, 409, 10, 35, 5, 35, 411, 10, 35,
	3, 36, 3, 36, 5, 36, 415, 10, 36, 3, 37, 3, 37, 3, 37, 3, 37, 3, 37, 3,
	38, 3, 38, 3, 38, 3, 38, 3, 39, 3, 39, 3, 39, 7, 39, 429, 10, 39, 12, 39,
	14, 39, 432, 11, 39, 3, 40, 3, 40, 3, 40, 5, 40, 437, 10, 40, 3, 40, 3,
	40, 3, 40, 3, 40, 5, 40, 443, 10, 40, 3, 40, 3, 40, 3, 40, 3, 40, 5, 40,
	449, 10, 40, 3, 40, 5, 40, 452, 10, 40, 3, 41, 3, 41, 3, 41, 3, 41, 3,
	41, 3, 41, 3, 41, 5, 41, 461, 10, 41, 3, 41, 3, 41, 3, 41, 3, 41, 3, 41,
	3, 41, 3, 41, 3, 41, 3, 41, 3, 41, 3, 41, 3, 41, 3, 41, 3, 41, 5, 41, 477,
	10, 41, 3, 41, 3, 41, 3, 41, 3, 41, 3, 41, 3, 41, 3, 41, 3, 41, 3, 41,
	3, 41, 3, 41, 3, 41, 3, 41, 3, 41, 3, 41, 3, 41, 3, 41, 3, 41, 3, 41, 3,
	41, 3, 41, 3, 41, 3, 41, 3, 41, 3, 41, 3, 41, 3, 41, 3, 41, 3, 41, 3, 41,
	3, 41, 3, 41, 3, 41, 3, 41, 3, 41, 3, 41, 5, 41, 515, 10, 41, 3, 41, 3,
	41, 3, 41, 3, 41, 3, 41, 3, 41, 3, 41, 3, 41, 3, 41, 3, 41, 7, 41, 527,
	10, 41, 12, 41, 14, 41, 530, 11, 41, 3, 42, 3, 42, 3, 42, 3, 42, 3, 42,
	3, 42, 3, 42, 3, 42, 5, 42, 540, 10, 42, 3, 43, 3, 43, 3, 43, 7, 43, 545,
	10, 43, 12, 43, 14, 43, 548, 11, 43, 3, 44, 3, 44, 3, 44, 5, 44, 553, 10,
	44, 3, 44, 3, 44, 7, 44, 557, 10, 44, 12, 44, 14, 44, 560, 11, 44, 3, 45,
	3, 45, 3, 45, 3, 45, 5, 45, 566, 10, 45, 3, 45, 3, 45, 3, 46, 3, 46, 3,
	47, 3, 47, 3, 47, 3, 48, 3, 48, 3, 48, 3, 48, 5, 48, 579, 10, 48, 5, 48,
	581, 10, 48, 3, 49, 3, 49, 5, 49, 585, 10, 49, 3, 49, 3, 49, 3, 49, 2,
	3, 80, 50, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22, 24, 26, 28, 30, 32,
	34, 36, 38, 40, 42, 44, 46, 48, 50, 52, 54, 56, 58, 60, 62, 64, 66, 68,
	70, 72, 74, 76, 78, 80, 82, 84, 86, 88, 90, 92, 94, 96, 2, 14, 4, 2, 19,
	19, 42, 42, 3, 2, 55, 58, 3, 2, 59, 60, 3, 2, 87, 90, 3, 2, 77, 78, 4,
	2, 91, 92, 96, 96, 3, 2, 89, 90, 4, 2, 75, 76, 82, 83, 4, 2, 81, 81, 84,
	84, 4, 2, 74, 74, 97, 107, 3, 2, 87, 88, 11, 2, 5, 5, 7, 7, 10, 10, 16,
	16, 22, 22, 29, 29, 31, 31, 39, 39, 54, 54, 2, 646, 2, 98, 3, 2, 2, 2,
	4, 110, 3, 2, 2, 2, 6, 121, 3, 2, 2, 2, 8, 126, 3, 2, 2, 2, 10, 129, 3,
	2, 2, 2, 12, 148, 3, 2, 2, 2, 14, 152, 3, 2, 2, 2, 16, 154, 3, 2, 2, 2,
	18, 162, 3, 2, 2, 2, 20, 181, 3, 2, 2, 2, 22, 186, 3, 2, 2, 2, 24, 195,
	3, 2, 2, 2, 26, 202, 3, 2, 2, 2, 28, 204, 3, 2, 2, 2, 30, 212, 3, 2, 2,
	2, 32, 215, 3, 2, 2, 2, 34, 223, 3, 2, 2, 2, 36, 232, 3, 2, 2, 2, 38, 237,
	3, 2, 2, 2, 40, 247, 3, 2, 2, 2, 42, 249, 3, 2, 2, 2, 44, 265, 3, 2, 2,
	2, 46, 279, 3, 2, 2, 2, 48, 287, 3, 2, 2, 2, 50, 289, 3, 2, 2, 2, 52, 291,
	3, 2, 2, 2, 54, 293, 3, 2, 2, 2, 56, 295, 3, 2, 2, 2, 58, 302, 3, 2, 2,
	2, 60, 311, 3, 2, 2, 2, 62, 376, 3, 2, 2, 2, 64, 379, 3, 2, 2, 2, 66, 396,
	3, 2, 2, 2, 68, 410, 3, 2, 2, 2, 70, 414, 3, 2, 2, 2, 72, 416, 3, 2, 2,
	2, 74, 421, 3, 2, 2, 2, 76, 425, 3, 2, 2, 2, 78, 451, 3, 2, 2, 2, 80, 460,
	3, 2, 2, 2, 82, 539, 3, 2, 2, 2, 84, 541, 3, 2, 2, 2, 86, 552, 3, 2, 2,
	2, 88, 561, 3, 2, 2, 2, 90, 569, 3, 2, 2, 2, 92, 571, 3, 2, 2, 2, 94, 580,
	3, 2, 2, 2, 96, 582, 3, 2, 2, 2, 98, 99, 7, 11, 2, 2, 99, 102, 7, 115,
	2, 2, 100, 101, 7, 19, 2, 2, 101, 103, 5, 86, 44, 2, 102, 100, 3, 2, 2,
	2, 102, 103, 3, 2, 2, 2, 103, 106, 3, 2, 2, 2, 104, 105, 7, 26, 2, 2, 105,
	107, 5, 84, 43, 2, 106, 104, 3, 2, 2, 2, 106, 107, 3, 2, 2, 2, 107, 108,
	3, 2, 2, 2, 108, 109, 5, 4, 3, 2, 109, 3, 3, 2, 2, 2, 110, 114, 7, 67,
	2, 2, 111, 113, 5, 6, 4, 2, 112, 111, 3, 2, 2, 2, 113, 116, 3, 2, 2, 2,
	114, 112, 3, 2, 2, 2, 114, 115, 3, 2, 2, 2, 115, 117, 3, 2, 2, 2, 116,
	114, 3, 2, 2, 2, 117, 118, 7, 68, 2, 2, 118, 5, 3, 2, 2, 2, 119, 122, 7,
	71, 2, 2, 120, 122, 5, 8, 5, 2, 121, 119, 3, 2, 2, 2, 121, 120, 3, 2, 2,
	2, 122, 7, 3, 2, 2, 2, 123, 127, 5, 10, 6, 2, 124, 127, 5, 30, 16, 2, 125,
	127, 5, 2, 2, 2, 126, 123, 3, 2, 2, 2, 126, 124, 3, 2, 2, 2, 126, 125,
	3, 2, 2, 2, 127, 9, 3, 2, 2, 2, 128, 130, 5, 14, 8, 2, 129, 128, 3, 2,
	2, 2, 129, 130, 3, 2, 2, 2, 130, 131, 3, 2, 2, 2, 131, 132, 7, 115, 2,
	2, 132, 137, 5, 18, 10, 2, 133, 134, 7, 69, 2, 2, 134, 136, 7, 70, 2, 2,
	135, 133, 3, 2, 2, 2, 136, 139, 3, 2, 2, 2, 137, 135, 3, 2, 2, 2, 137,
	138, 3, 2, 2, 2, 138, 142, 3, 2, 2, 2, 139, 137, 3, 2, 2, 2, 140, 141,
	7, 47, 2, 2, 141, 143, 5, 16, 9, 2, 142, 140, 3, 2, 2, 2, 142, 143, 3,
	2, 2, 2, 143, 144, 3, 2, 2, 2, 144, 145, 5, 12, 7, 2, 145, 11, 3, 2, 2,
	2, 146, 149, 5, 56, 29, 2, 147, 149, 7, 71, 2, 2, 148, 146, 3, 2, 2, 2,
	148, 147, 3, 2, 2, 2, 149, 13, 3, 2, 2, 2, 150, 153, 5, 86, 44, 2, 151,
	153, 7, 50, 2, 2, 152, 150, 3, 2, 2, 2, 152, 151, 3, 2, 2, 2, 153, 15,
	3, 2, 2, 2, 154, 159, 5, 28, 15, 2, 155, 156, 7, 72, 2, 2, 156, 158, 5,
	28, 15, 2, 157, 155, 3, 2, 2, 2, 158, 161, 3, 2, 2, 2, 159, 157, 3, 2,
	2, 2, 159, 160, 3, 2, 2, 2, 160, 17, 3, 2, 2, 2, 161, 159, 3, 2, 2, 2,
	162, 164, 7, 65, 2, 2, 163, 165, 5, 20, 11, 2, 164, 163, 3, 2, 2, 2, 164,
	165, 3, 2, 2, 2, 165, 166, 3, 2, 2, 2, 166, 167, 7, 66, 2, 2, 167, 19,
	3, 2, 2, 2, 168, 173, 5, 22, 12, 2, 169, 170, 7, 72, 2, 2, 170, 172, 5,
	22, 12, 2, 171, 169, 3, 2, 2, 2, 172, 175, 3, 2, 2, 2, 173, 171, 3, 2,
	2, 2, 173, 174, 3, 2, 2, 2, 174, 178, 3, 2, 2, 2, 175, 173, 3, 2, 2, 2,
	176, 177, 7, 72, 2, 2, 177, 179, 5, 24, 13, 2, 178, 176, 3, 2, 2, 2, 178,
	179, 3, 2, 2, 2, 179, 182, 3, 2, 2, 2, 180, 182, 5, 24, 13, 2, 181, 168,
	3, 2, 2, 2, 181, 180, 3, 2, 2, 2, 182, 21, 3, 2, 2, 2, 183, 185, 5, 26,
	14, 2, 184, 183, 3, 2, 2, 2, 185, 188, 3, 2, 2, 2, 186, 184, 3, 2, 2, 2,
	186, 187, 3, 2, 2, 2, 187, 189, 3, 2, 2, 2, 188, 186, 3, 2, 2, 2, 189,
	190, 5, 86, 44, 2, 190, 191, 5, 38, 20, 2, 191, 23, 3, 2, 2, 2, 192, 194,
	5, 26, 14, 2, 193, 192, 3, 2, 2, 2, 194, 197, 3, 2, 2, 2, 195, 193, 3,
	2, 2, 2, 195, 196, 3, 2, 2, 2, 196, 198, 3, 2, 2, 2, 197, 195, 3, 2, 2,
	2, 198, 199, 5, 86, 44, 2, 199, 200, 7, 111, 2, 2, 200, 201, 5, 38, 20,
	2, 201, 25, 3, 2, 2, 2, 202, 203, 7, 20, 2, 2, 203, 27, 3, 2, 2, 2, 204,
	209, 7, 115, 2, 2, 205, 206, 7, 73, 2, 2, 206, 208, 7, 115, 2, 2, 207,
	205, 3, 2, 2, 2, 208, 211, 3, 2, 2, 2, 209, 207, 3, 2, 2, 2, 209, 210,
	3, 2, 2, 2, 210, 29, 3, 2, 2, 2, 211, 209, 3, 2, 2, 2, 212, 213, 5, 34,
	18, 2, 213, 214, 7, 71, 2, 2, 214, 31, 3, 2, 2, 2, 215, 216, 7, 115, 2,
	2, 216, 219, 5, 18, 10, 2, 217, 218, 7, 47, 2, 2, 218, 220, 5, 16, 9, 2,
	219, 217, 3, 2, 2, 2, 219, 220, 3, 2, 2, 2, 220, 221, 3, 2, 2, 2, 221,
	222, 5, 56, 29, 2, 222, 33, 3, 2, 2, 2, 223, 224, 5, 86, 44, 2, 224, 229,
	5, 36, 19, 2, 225, 226, 7, 72, 2, 2, 226, 228, 5, 36, 19, 2, 227, 225,
	3, 2, 2, 2, 228, 231, 3, 2, 2, 2, 229, 227, 3, 2, 2, 2, 229, 230, 3, 2,
	2, 2, 230, 35, 3, 2, 2, 2, 231, 229, 3, 2, 2, 2, 232, 235, 5, 38, 20, 2,
	233, 234, 7, 74, 2, 2, 234, 236, 5, 40, 21, 2, 235, 233, 3, 2, 2, 2, 235,
	236, 3, 2, 2, 2, 236, 37, 3, 2, 2, 2, 237, 242, 7, 115, 2, 2, 238, 239,
	7, 69, 2, 2, 239, 241, 7, 70, 2, 2, 240, 238, 3, 2, 2, 2, 241, 244, 3,
	2, 2, 2, 242, 240, 3, 2, 2, 2, 242, 243, 3, 2, 2, 2, 243, 39, 3, 2, 2,
	2, 244, 242, 3, 2, 2, 2, 245, 248, 5, 42, 22, 2, 246, 248, 5, 80, 41, 2,
	247, 245, 3, 2, 2, 2, 247, 246, 3, 2, 2, 2, 248, 41, 3, 2, 2, 2, 249, 261,
	7, 67, 2, 2, 250, 255, 5, 40, 21, 2, 251, 252, 7, 72, 2, 2, 252, 254, 5,
	40, 21, 2, 253, 251, 3, 2, 2, 2, 254, 257, 3, 2, 2, 2, 255, 253, 3, 2,
	2, 2, 255, 256, 3, 2, 2, 2, 256, 259, 3, 2, 2, 2, 257, 255, 3, 2, 2, 2,
	258, 260, 7, 72, 2, 2, 259, 258, 3, 2, 2, 2, 259, 260, 3, 2, 2, 2, 260,
	262, 3, 2, 2, 2, 261, 250, 3, 2, 2, 2, 261, 262, 3, 2, 2, 2, 262, 263,
	3, 2, 2, 2, 263, 264, 7, 68, 2, 2, 264, 43, 3, 2, 2, 2, 265, 270, 7, 115,
	2, 2, 266, 267, 7, 73, 2, 2, 267, 269, 7, 115, 2, 2, 268, 266, 3, 2, 2,
	2, 269, 272, 3, 2, 2, 2, 270, 268, 3, 2, 2, 2, 270, 271, 3, 2, 2, 2, 271,
	45, 3, 2, 2, 2, 272, 270, 3, 2, 2, 2, 273, 280, 5, 86, 44, 2, 274, 277,
	7, 79, 2, 2, 275, 276, 9, 2, 2, 2, 276, 278, 5, 86, 44, 2, 277, 275, 3,
	2, 2, 2, 277, 278, 3, 2, 2, 2, 278, 280, 3, 2, 2, 2, 279, 273, 3, 2, 2,
	2, 279, 274, 3, 2, 2, 2, 280, 47, 3, 2, 2, 2, 281, 288, 5, 50, 26, 2, 282,
	288, 5, 52, 27, 2, 283, 288, 7, 62, 2, 2, 284, 288, 7, 63, 2, 2, 285, 288,
	7, 61, 2, 2, 286, 288, 7, 64, 2, 2, 287, 281, 3, 2, 2, 2, 287, 282, 3,
	2, 2, 2, 287, 283, 3, 2, 2, 2, 287, 284, 3, 2, 2, 2, 287, 285, 3, 2, 2,
	2, 287, 286, 3, 2, 2, 2, 288, 49, 3, 2, 2, 2, 289, 290, 9, 3, 2, 2, 290,
	51, 3, 2, 2, 2, 291, 292, 9, 4, 2, 2, 292, 53, 3, 2, 2, 2, 293, 294, 5,
	58, 30, 2, 294, 55, 3, 2, 2, 2, 295, 296, 7, 67, 2, 2, 296, 297, 5, 58,
	30, 2, 297, 298, 7, 68, 2, 2, 298, 57, 3, 2, 2, 2, 299, 301, 5, 60, 31,
	2, 300, 299, 3, 2, 2, 2, 301, 304, 3, 2, 2, 2, 302, 300, 3, 2, 2, 2, 302,
	303, 3, 2, 2, 2, 303, 59, 3, 2, 2, 2, 304, 302, 3, 2, 2, 2, 305, 306, 5,
	34, 18, 2, 306, 307, 7, 71, 2, 2, 307, 312, 3, 2, 2, 2, 308, 312, 5, 62,
	32, 2, 309, 312, 5, 10, 6, 2, 310, 312, 5, 2, 2, 2, 311, 305, 3, 2, 2,
	2, 311, 308, 3, 2, 2, 2, 311, 309, 3, 2, 2, 2, 311, 310, 3, 2, 2, 2, 312,
	61, 3, 2, 2, 2, 313, 377, 5, 56, 29, 2, 314, 315, 7, 24, 2, 2, 315, 316,
	5, 74, 38, 2, 316, 319, 5, 62, 32, 2, 317, 318, 7, 17, 2, 2, 318, 320,
	5, 62, 32, 2, 319, 317, 3, 2, 2, 2, 319, 320, 3, 2, 2, 2, 320, 377, 3,
	2, 2, 2, 321, 322, 7, 23, 2, 2, 322, 323, 7, 65, 2, 2, 323, 324, 5, 68,
	35, 2, 324, 325, 7, 66, 2, 2, 325, 326, 5, 62, 32, 2, 326, 377, 3, 2, 2,
	2, 327, 328, 7, 52, 2, 2, 328, 329, 5, 74, 38, 2, 329, 330, 5, 62, 32,
	2, 330, 377, 3, 2, 2, 2, 331, 332, 7, 15, 2, 2, 332, 333, 5, 62, 32, 2,
	333, 334, 7, 52, 2, 2, 334, 335, 5, 74, 38, 2, 335, 336, 7, 71, 2, 2, 336,
	377, 3, 2, 2, 2, 337, 338, 7, 43, 2, 2, 338, 339, 5, 74, 38, 2, 339, 343,
	7, 67, 2, 2, 340, 342, 5, 64, 33, 2, 341, 340, 3, 2, 2, 2, 342, 345, 3,
	2, 2, 2, 343, 341, 3, 2, 2, 2, 343, 344, 3, 2, 2, 2, 344, 349, 3, 2, 2,
	2, 345, 343, 3, 2, 2, 2, 346, 348, 5, 66, 34, 2, 347, 346, 3, 2, 2, 2,
	348, 351, 3, 2, 2, 2, 349, 347, 3, 2, 2, 2, 349, 350, 3, 2, 2, 2, 350,
	352, 3, 2, 2, 2, 351, 349, 3, 2, 2, 2, 352, 353, 7, 68, 2, 2, 353, 377,
	3, 2, 2, 2, 354, 356, 7, 38, 2, 2, 355, 357, 5, 80, 41, 2, 356, 355, 3,
	2, 2, 2, 356, 357, 3, 2, 2, 2, 357, 358, 3, 2, 2, 2, 358, 377, 7, 71, 2,
	2, 359, 361, 7, 6, 2, 2, 360, 362, 7, 115, 2, 2, 361, 360, 3, 2, 2, 2,
	361, 362, 3, 2, 2, 2, 362, 363, 3, 2, 2, 2, 363, 377, 7, 71, 2, 2, 364,
	366, 7, 13, 2, 2, 365, 367, 7, 115, 2, 2, 366, 365, 3, 2, 2, 2, 366, 367,
	3, 2, 2, 2, 367, 368, 3, 2, 2, 2, 368, 377, 7, 71, 2, 2, 369, 377, 7, 71,
	2, 2, 370, 371, 5, 80, 41, 2, 371, 372, 7, 71, 2, 2, 372, 377, 3, 2, 2,
	2, 373, 374, 7, 115, 2, 2, 374, 375, 7, 80, 2, 2, 375, 377, 5, 62, 32,
	2, 376, 313, 3, 2, 2, 2, 376, 314, 3, 2, 2, 2, 376, 321, 3, 2, 2, 2, 376,
	327, 3, 2, 2, 2, 376, 331, 3, 2, 2, 2, 376, 337, 3, 2, 2, 2, 376, 354,
	3, 2, 2, 2, 376, 359, 3, 2, 2, 2, 376, 364, 3, 2, 2, 2, 376, 369, 3, 2,
	2, 2, 376, 370, 3, 2, 2, 2, 376, 373, 3, 2, 2, 2, 377, 63, 3, 2, 2, 2,
	378, 380, 5, 66, 34, 2, 379, 378, 3, 2, 2, 2, 380, 381, 3, 2, 2, 2, 381,
	379, 3, 2, 2, 2, 381, 382, 3, 2, 2, 2, 382, 384, 3, 2, 2, 2, 383, 385,
	5, 60, 31, 2, 384, 383, 3, 2, 2, 2, 385, 386, 3, 2, 2, 2, 386, 384, 3,
	2, 2, 2, 386, 387, 3, 2, 2, 2, 387, 65, 3, 2, 2, 2, 388, 391, 7, 8, 2,
	2, 389, 392, 5, 80, 41, 2, 390, 392, 7, 115, 2, 2, 391, 389, 3, 2, 2, 2,
	391, 390, 3, 2, 2, 2, 392, 393, 3, 2, 2, 2, 393, 397, 7, 80, 2, 2, 394,
	395, 7, 14, 2, 2, 395, 397, 7, 80, 2, 2, 396, 388, 3, 2, 2, 2, 396, 394,
	3, 2, 2, 2, 397, 67, 3, 2, 2, 2, 398, 411, 5, 72, 37, 2, 399, 401, 5, 70,
	36, 2, 400, 399, 3, 2, 2, 2, 400, 401, 3, 2, 2, 2, 401, 402, 3, 2, 2, 2,
	402, 404, 7, 71, 2, 2, 403, 405, 5, 80, 41, 2, 404, 403, 3, 2, 2, 2, 404,
	405, 3, 2, 2, 2, 405, 406, 3, 2, 2, 2, 406, 408, 7, 71, 2, 2, 407, 409,
	5, 76, 39, 2, 408, 407, 3, 2, 2, 2, 408, 409, 3, 2, 2, 2, 409, 411, 3,
	2, 2, 2, 410, 398, 3, 2, 2, 2, 410, 400, 3, 2, 2, 2, 411, 69, 3, 2, 2,
	2, 412, 415, 5, 34, 18, 2, 413, 415, 5, 76, 39, 2, 414, 412, 3, 2, 2, 2,
	414, 413, 3, 2, 2, 2, 415, 71, 3, 2, 2, 2, 416, 417, 5, 86, 44, 2, 417,
	418, 5, 38, 20, 2, 418, 419, 7, 80, 2, 2, 419, 420, 5, 80, 41, 2, 420,
	73, 3, 2, 2, 2, 421, 422, 7, 65, 2, 2, 422, 423, 5, 80, 41, 2, 423, 424,
	7, 66, 2, 2, 424, 75, 3, 2, 2, 2, 425, 430, 5, 80, 41, 2, 426, 427, 7,
	72, 2, 2, 427, 429, 5, 80, 41, 2, 428, 426, 3, 2, 2, 2, 429, 432, 3, 2,
	2, 2, 430, 428, 3, 2, 2, 2, 430, 431, 3, 2, 2, 2, 431, 77, 3, 2, 2, 2,
	432, 430, 3, 2, 2, 2, 433, 434, 7, 115, 2, 2, 434, 436, 7, 65, 2, 2, 435,
	437, 5, 76, 39, 2, 436, 435, 3, 2, 2, 2, 436, 437, 3, 2, 2, 2, 437, 438,
	3, 2, 2, 2, 438, 452, 7, 66, 2, 2, 439, 440, 7, 45, 2, 2, 440, 442, 7,
	65, 2, 2, 441, 443, 5, 76, 39, 2, 442, 441, 3, 2, 2, 2, 442, 443, 3, 2,
	2, 2, 443, 444, 3, 2, 2, 2, 444, 452, 7, 66, 2, 2, 445, 446, 7, 42, 2,
	2, 446, 448, 7, 65, 2, 2, 447, 449, 5, 76, 39, 2, 448, 447, 3, 2, 2, 2,
	448, 449, 3, 2, 2, 2, 449, 450, 3, 2, 2, 2, 450, 452, 7, 66, 2, 2, 451,
	433, 3, 2, 2, 2, 451, 439, 3, 2, 2, 2, 451, 445, 3, 2, 2, 2, 452, 79, 3,
	2, 2, 2, 453, 454, 8, 41, 1, 2, 454, 461, 5, 82, 42, 2, 455, 461, 5, 78,
	40, 2, 456, 457, 9, 5, 2, 2, 457, 461, 5, 80, 41, 17, 458, 459, 9, 6, 2,
	2, 459, 461, 5, 80, 41, 16, 460, 453, 3, 2, 2, 2, 460, 455, 3, 2, 2, 2,
	460, 456, 3, 2, 2, 2, 460, 458, 3, 2, 2, 2, 461, 528, 3, 2, 2, 2, 462,
	463, 12, 15, 2, 2, 463, 464, 9, 7, 2, 2, 464, 527, 5, 80, 41, 16, 465,
	466, 12, 14, 2, 2, 466, 467, 9, 8, 2, 2, 467, 527, 5, 80, 41, 15, 468,
	476, 12, 13, 2, 2, 469, 470, 7, 76, 2, 2, 470, 477, 7, 76, 2, 2, 471, 472,
	7, 75, 2, 2, 472, 473, 7, 75, 2, 2, 473, 477, 7, 75, 2, 2, 474, 475, 7,
	75, 2, 2, 475, 477, 7, 75, 2, 2, 476, 469, 3, 2, 2, 2, 476, 471, 3, 2,
	2, 2, 476, 474, 3, 2, 2, 2, 477, 478, 3, 2, 2, 2, 478, 527, 5, 80, 41,
	14, 479, 480, 12, 12, 2, 2, 480, 481, 9, 9, 2, 2, 481, 527, 5, 80, 41,
	13, 482, 483, 12, 10, 2, 2, 483, 484, 9, 10, 2, 2, 484, 527, 5, 80, 41,
	11, 485, 486, 12, 9, 2, 2, 486, 487, 7, 93, 2, 2, 487, 527, 5, 80, 41,
	10, 488, 489, 12, 8, 2, 2, 489, 490, 7, 95, 2, 2, 490, 527, 5, 80, 41,
	9, 491, 492, 12, 7, 2, 2, 492, 493, 7, 94, 2, 2, 493, 527, 5, 80, 41, 8,
	494, 495, 12, 6, 2, 2, 495, 496, 7, 85, 2, 2, 496, 527, 5, 80, 41, 7, 497,
	498, 12, 5, 2, 2, 498, 499, 7, 86, 2, 2, 499, 527, 5, 80, 41, 6, 500, 501,
	12, 4, 2, 2, 501, 502, 7, 79, 2, 2, 502, 503, 5, 80, 41, 2, 503, 504, 7,
	80, 2, 2, 504, 505, 5, 80, 41, 5, 505, 527, 3, 2, 2, 2, 506, 507, 12, 3,
	2, 2, 507, 508, 9, 11, 2, 2, 508, 527, 5, 80, 41, 3, 509, 510, 12, 21,
	2, 2, 510, 514, 7, 73, 2, 2, 511, 515, 7, 115, 2, 2, 512, 515, 5, 78, 40,
	2, 513, 515, 7, 45, 2, 2, 514, 511, 3, 2, 2, 2, 514, 512, 3, 2, 2, 2, 514,
	513, 3, 2, 2, 2, 515, 527, 3, 2, 2, 2, 516, 517, 12, 20, 2, 2, 517, 518,
	7, 69, 2, 2, 518, 519, 5, 80, 41, 2, 519, 520, 7, 70, 2, 2, 520, 527, 3,
	2, 2, 2, 521, 522, 12, 18, 2, 2, 522, 527, 9, 12, 2, 2, 523, 524, 12, 11,
	2, 2, 524, 525, 7, 28, 2, 2, 525, 527, 5, 86, 44, 2, 526, 462, 3, 2, 2,
	2, 526, 465, 3, 2, 2, 2, 526, 468, 3, 2, 2, 2, 526, 479, 3, 2, 2, 2, 526,
	482, 3, 2, 2, 2, 526, 485, 3, 2, 2, 2, 526, 488, 3, 2, 2, 2, 526, 491,
	3, 2, 2, 2, 526, 494, 3, 2, 2, 2, 526, 497, 3, 2, 2, 2, 526, 500, 3, 2,
	2, 2, 526, 506, 3, 2, 2, 2, 526, 509, 3, 2, 2, 2, 526, 516, 3, 2, 2, 2,
	526, 521, 3, 2, 2, 2, 526, 523, 3, 2, 2, 2, 527, 530, 3, 2, 2, 2, 528,
	526, 3, 2, 2, 2, 528, 529, 3, 2, 2, 2, 529, 81, 3, 2, 2, 2, 530, 528, 3,
	2, 2, 2, 531, 532, 7, 65, 2, 2, 532, 533, 5, 80, 41, 2, 533, 534, 7, 66,
	2, 2, 534, 540, 3, 2, 2, 2, 535, 540, 7, 45, 2, 2, 536, 540, 7, 42, 2,
	2, 537, 540, 5, 48, 25, 2, 538, 540, 7, 115, 2, 2, 539, 531, 3, 2, 2, 2,
	539, 535, 3, 2, 2, 2, 539, 536, 3, 2, 2, 2, 539, 537, 3, 2, 2, 2, 539,
	538, 3, 2, 2, 2, 540, 83, 3, 2, 2, 2, 541, 546, 5, 86, 44, 2, 542, 543,
	7, 72, 2, 2, 543, 545, 5, 86, 44, 2, 544, 542, 3, 2, 2, 2, 545, 548, 3,
	2, 2, 2, 546, 544, 3, 2, 2, 2, 546, 547, 3, 2, 2, 2, 547, 85, 3, 2, 2,
	2, 548, 546, 3, 2, 2, 2, 549, 553, 5, 44, 23, 2, 550, 553, 5, 88, 45, 2,
	551, 553, 5, 90, 46, 2, 552, 549, 3, 2, 2, 2, 552, 550, 3, 2, 2, 2, 552,
	551, 3, 2, 2, 2, 553, 558, 3, 2, 2, 2, 554, 555, 7, 69, 2, 2, 555, 557,
	7, 70, 2, 2, 556, 554, 3, 2, 2, 2, 557, 560, 3, 2, 2, 2, 558, 556, 3, 2,
	2, 2, 558, 559, 3, 2, 2, 2, 559, 87, 3, 2, 2, 2, 560, 558, 3, 2, 2, 2,
	561, 562, 7, 53, 2, 2, 562, 563, 5, 14, 8, 2, 563, 565, 7, 65, 2, 2, 564,
	566, 5, 84, 43, 2, 565, 564, 3, 2, 2, 2, 565, 566, 3, 2, 2, 2, 566, 567,
	3, 2, 2, 2, 567, 568, 7, 66, 2, 2, 568, 89, 3, 2, 2, 2, 569, 570, 9, 13,
	2, 2, 570, 91, 3, 2, 2, 2, 571, 572, 7, 115, 2, 2, 572, 573, 5, 96, 49,
	2, 573, 93, 3, 2, 2, 2, 574, 581, 5, 96, 49, 2, 575, 576, 7, 73, 2, 2,
	576, 578, 7, 115, 2, 2, 577, 579, 5, 96, 49, 2, 578, 577, 3, 2, 2, 2, 578,
	579, 3, 2, 2, 2, 579, 581, 3, 2, 2, 2, 580, 574, 3, 2, 2, 2, 580, 575,
	3, 2, 2, 2, 581, 95, 3, 2, 2, 2, 582, 584, 7, 65, 2, 2, 583, 585, 5, 76,
	39, 2, 584, 583, 3, 2, 2, 2, 584, 585, 3, 2, 2, 2, 585, 586, 3, 2, 2, 2,
	586, 587, 7, 66, 2, 2, 587, 97, 3, 2, 2, 2, 68, 102, 106, 114, 121, 126,
	129, 137, 142, 148, 152, 159, 164, 173, 178, 181, 186, 195, 209, 219, 229,
	235, 242, 247, 255, 259, 261, 270, 277, 279, 287, 302, 311, 319, 343, 349,
	356, 361, 366, 376, 381, 386, 391, 396, 400, 404, 408, 410, 414, 430, 436,
	442, 448, 451, 460, 476, 514, 526, 528, 539, 546, 552, 558, 565, 578, 580,
	584,
}
var deserializer = antlr.NewATNDeserializer(nil)
var deserializedATN = deserializer.DeserializeFromUInt16(parserATN)

var literalNames = []string{
	"", "'abstract'", "'assert'", "'boolean'", "'break'", "'byte'", "'case'",
	"'catch'", "'char'", "'class'", "'const'", "'continue'", "'default'", "'do'",
	"'double'", "'else'", "'enum'", "'extends'", "'final'", "'finally'", "'float'",
	"'for'", "'if'", "'goto'", "'implements'", "'import'", "'instanceof'",
	"'int'", "'interface'", "'long'", "'native'", "'new'", "'package'", "'private'",
	"'protected'", "'public'", "'return'", "'short'", "'static'", "'strictfp'",
	"'super'", "'switch'", "'synchronized'", "'this'", "'throw'", "'throws'",
	"'transient'", "'try'", "'void'", "'volatile'", "'while'", "'function'",
	"'string'", "", "", "", "", "", "", "", "", "", "'null'", "'('", "')'",
	"'{'", "'}'", "'['", "']'", "';'", "','", "'.'", "'='", "'>'", "'<'", "'!'",
	"'~'", "'?'", "':'", "'=='", "'<='", "'>='", "'!='", "'&&'", "'||'", "'++'",
	"'--'", "'+'", "'-'", "'*'", "'/'", "'&'", "'|'", "'^'", "'%'", "'+='",
	"'-='", "'*='", "'/='", "'&='", "'|='", "'^='", "'%='", "'<<='", "'>>='",
	"'>>>='", "'->'", "'::'", "'@'", "'...'",
}
var symbolicNames = []string{
	"", "ABSTRACT", "ASSERT", "BOOLEAN", "BREAK", "BYTE", "CASE", "CATCH",
	"CHAR", "CLASS", "CONST", "CONTINUE", "DEFAULT", "DO", "DOUBLE", "ELSE",
	"ENUM", "EXTENDS", "FINAL", "FINALLY", "FLOAT", "FOR", "IF", "GOTO", "IMPLEMENTS",
	"IMPORT", "INSTANCEOF", "INT", "INTERFACE", "LONG", "NATIVE", "NEW", "PACKAGE",
	"PRIVATE", "PROTECTED", "PUBLIC", "RETURN", "SHORT", "STATIC", "STRICTFP",
	"SUPER", "SWITCH", "SYNCHRONIZED", "THIS", "THROW", "THROWS", "TRANSIENT",
	"TRY", "VOID", "VOLATILE", "WHILE", "FUNCTION", "STRING", "DECIMAL_LITERAL",
	"HEX_LITERAL", "OCT_LITERAL", "BINARY_LITERAL", "FLOAT_LITERAL", "HEX_FLOAT_LITERAL",
	"BOOL_LITERAL", "CHAR_LITERAL", "STRING_LITERAL", "NULL_LITERAL", "LPAREN",
	"RPAREN", "LBRACE", "RBRACE", "LBRACK", "RBRACK", "SEMI", "COMMA", "DOT",
	"ASSIGN", "GT", "LT", "BANG", "TILDE", "QUESTION", "COLON", "EQUAL", "LE",
	"GE", "NOTEQUAL", "AND", "OR", "INC", "DEC", "ADD", "SUB", "MUL", "DIV",
	"BITAND", "BITOR", "CARET", "MOD", "ADD_ASSIGN", "SUB_ASSIGN", "MUL_ASSIGN",
	"DIV_ASSIGN", "AND_ASSIGN", "OR_ASSIGN", "XOR_ASSIGN", "MOD_ASSIGN", "LSHIFT_ASSIGN",
	"RSHIFT_ASSIGN", "URSHIFT_ASSIGN", "ARROW", "COLONCOLON", "AT", "ELLIPSIS",
	"WS", "COMMENT", "LINE_COMMENT", "IDENTIFIER",
}

var ruleNames = []string{
	"classDeclaration", "classBody", "classBodyDeclaration", "memberDeclaration",
	"functionDeclaration", "functionBody", "typeTypeOrVoid", "qualifiedNameList",
	"formalParameters", "formalParameterList", "formalParameter", "lastFormalParameter",
	"variableModifier", "qualifiedName", "fieldDeclaration", "constructorDeclaration",
	"variableDeclarators", "variableDeclarator", "variableDeclaratorId", "variableInitializer",
	"arrayInitializer", "classOrInterfaceType", "typeArgument", "literal",
	"integerLiteral", "floatLiteral", "prog", "block", "blockStatements", "blockStatement",
	"statement", "switchBlockStatementGroup", "switchLabel", "forControl",
	"forInit", "enhancedForControl", "parExpression", "expressionList", "functionCall",
	"expression", "primary", "typeList", "typeType", "functionType", "primitiveType",
	"creator", "superSuffix", "arguments",
}
var decisionToDFA = make([]*antlr.DFA, len(deserializedATN.DecisionToState))

func init() {
	for index, ds := range deserializedATN.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(ds, index)
	}
}

type PlayScriptParser struct {
	*antlr.BaseParser
}

func NewPlayScriptParser(input antlr.TokenStream) *PlayScriptParser {
	this := new(PlayScriptParser)

	this.BaseParser = antlr.NewBaseParser(input)

	this.Interpreter = antlr.NewParserATNSimulator(this, deserializedATN, decisionToDFA, antlr.NewPredictionContextCache())
	this.RuleNames = ruleNames
	this.LiteralNames = literalNames
	this.SymbolicNames = symbolicNames
	this.GrammarFileName = "PlayScript.g4"

	return this
}

// PlayScriptParser tokens.
const (
	PlayScriptParserEOF               = antlr.TokenEOF
	PlayScriptParserABSTRACT          = 1
	PlayScriptParserASSERT            = 2
	PlayScriptParserBOOLEAN           = 3
	PlayScriptParserBREAK             = 4
	PlayScriptParserBYTE              = 5
	PlayScriptParserCASE              = 6
	PlayScriptParserCATCH             = 7
	PlayScriptParserCHAR              = 8
	PlayScriptParserCLASS             = 9
	PlayScriptParserCONST             = 10
	PlayScriptParserCONTINUE          = 11
	PlayScriptParserDEFAULT           = 12
	PlayScriptParserDO                = 13
	PlayScriptParserDOUBLE            = 14
	PlayScriptParserELSE              = 15
	PlayScriptParserENUM              = 16
	PlayScriptParserEXTENDS           = 17
	PlayScriptParserFINAL             = 18
	PlayScriptParserFINALLY           = 19
	PlayScriptParserFLOAT             = 20
	PlayScriptParserFOR               = 21
	PlayScriptParserIF                = 22
	PlayScriptParserGOTO              = 23
	PlayScriptParserIMPLEMENTS        = 24
	PlayScriptParserIMPORT            = 25
	PlayScriptParserINSTANCEOF        = 26
	PlayScriptParserINT               = 27
	PlayScriptParserINTERFACE         = 28
	PlayScriptParserLONG              = 29
	PlayScriptParserNATIVE            = 30
	PlayScriptParserNEW               = 31
	PlayScriptParserPACKAGE           = 32
	PlayScriptParserPRIVATE           = 33
	PlayScriptParserPROTECTED         = 34
	PlayScriptParserPUBLIC            = 35
	PlayScriptParserRETURN            = 36
	PlayScriptParserSHORT             = 37
	PlayScriptParserSTATIC            = 38
	PlayScriptParserSTRICTFP          = 39
	PlayScriptParserSUPER             = 40
	PlayScriptParserSWITCH            = 41
	PlayScriptParserSYNCHRONIZED      = 42
	PlayScriptParserTHIS              = 43
	PlayScriptParserTHROW             = 44
	PlayScriptParserTHROWS            = 45
	PlayScriptParserTRANSIENT         = 46
	PlayScriptParserTRY               = 47
	PlayScriptParserVOID              = 48
	PlayScriptParserVOLATILE          = 49
	PlayScriptParserWHILE             = 50
	PlayScriptParserFUNCTION          = 51
	PlayScriptParserSTRING            = 52
	PlayScriptParserDECIMAL_LITERAL   = 53
	PlayScriptParserHEX_LITERAL       = 54
	PlayScriptParserOCT_LITERAL       = 55
	PlayScriptParserBINARY_LITERAL    = 56
	PlayScriptParserFLOAT_LITERAL     = 57
	PlayScriptParserHEX_FLOAT_LITERAL = 58
	PlayScriptParserBOOL_LITERAL      = 59
	PlayScriptParserCHAR_LITERAL      = 60
	PlayScriptParserSTRING_LITERAL    = 61
	PlayScriptParserNULL_LITERAL      = 62
	PlayScriptParserLPAREN            = 63
	PlayScriptParserRPAREN            = 64
	PlayScriptParserLBRACE            = 65
	PlayScriptParserRBRACE            = 66
	PlayScriptParserLBRACK            = 67
	PlayScriptParserRBRACK            = 68
	PlayScriptParserSEMI              = 69
	PlayScriptParserCOMMA             = 70
	PlayScriptParserDOT               = 71
	PlayScriptParserASSIGN            = 72
	PlayScriptParserGT                = 73
	PlayScriptParserLT                = 74
	PlayScriptParserBANG              = 75
	PlayScriptParserTILDE             = 76
	PlayScriptParserQUESTION          = 77
	PlayScriptParserCOLON             = 78
	PlayScriptParserEQUAL             = 79
	PlayScriptParserLE                = 80
	PlayScriptParserGE                = 81
	PlayScriptParserNOTEQUAL          = 82
	PlayScriptParserAND               = 83
	PlayScriptParserOR                = 84
	PlayScriptParserINC               = 85
	PlayScriptParserDEC               = 86
	PlayScriptParserADD               = 87
	PlayScriptParserSUB               = 88
	PlayScriptParserMUL               = 89
	PlayScriptParserDIV               = 90
	PlayScriptParserBITAND            = 91
	PlayScriptParserBITOR             = 92
	PlayScriptParserCARET             = 93
	PlayScriptParserMOD               = 94
	PlayScriptParserADD_ASSIGN        = 95
	PlayScriptParserSUB_ASSIGN        = 96
	PlayScriptParserMUL_ASSIGN        = 97
	PlayScriptParserDIV_ASSIGN        = 98
	PlayScriptParserAND_ASSIGN        = 99
	PlayScriptParserOR_ASSIGN         = 100
	PlayScriptParserXOR_ASSIGN        = 101
	PlayScriptParserMOD_ASSIGN        = 102
	PlayScriptParserLSHIFT_ASSIGN     = 103
	PlayScriptParserRSHIFT_ASSIGN     = 104
	PlayScriptParserURSHIFT_ASSIGN    = 105
	PlayScriptParserARROW             = 106
	PlayScriptParserCOLONCOLON        = 107
	PlayScriptParserAT                = 108
	PlayScriptParserELLIPSIS          = 109
	PlayScriptParserWS                = 110
	PlayScriptParserCOMMENT           = 111
	PlayScriptParserLINE_COMMENT      = 112
	PlayScriptParserIDENTIFIER        = 113
)

// PlayScriptParser rules.
const (
	PlayScriptParserRULE_classDeclaration          = 0
	PlayScriptParserRULE_classBody                 = 1
	PlayScriptParserRULE_classBodyDeclaration      = 2
	PlayScriptParserRULE_memberDeclaration         = 3
	PlayScriptParserRULE_functionDeclaration       = 4
	PlayScriptParserRULE_functionBody              = 5
	PlayScriptParserRULE_typeTypeOrVoid            = 6
	PlayScriptParserRULE_qualifiedNameList         = 7
	PlayScriptParserRULE_formalParameters          = 8
	PlayScriptParserRULE_formalParameterList       = 9
	PlayScriptParserRULE_formalParameter           = 10
	PlayScriptParserRULE_lastFormalParameter       = 11
	PlayScriptParserRULE_variableModifier          = 12
	PlayScriptParserRULE_qualifiedName             = 13
	PlayScriptParserRULE_fieldDeclaration          = 14
	PlayScriptParserRULE_constructorDeclaration    = 15
	PlayScriptParserRULE_variableDeclarators       = 16
	PlayScriptParserRULE_variableDeclarator        = 17
	PlayScriptParserRULE_variableDeclaratorId      = 18
	PlayScriptParserRULE_variableInitializer       = 19
	PlayScriptParserRULE_arrayInitializer          = 20
	PlayScriptParserRULE_classOrInterfaceType      = 21
	PlayScriptParserRULE_typeArgument              = 22
	PlayScriptParserRULE_literal                   = 23
	PlayScriptParserRULE_integerLiteral            = 24
	PlayScriptParserRULE_floatLiteral              = 25
	PlayScriptParserRULE_prog                      = 26
	PlayScriptParserRULE_block                     = 27
	PlayScriptParserRULE_blockStatements           = 28
	PlayScriptParserRULE_blockStatement            = 29
	PlayScriptParserRULE_statement                 = 30
	PlayScriptParserRULE_switchBlockStatementGroup = 31
	PlayScriptParserRULE_switchLabel               = 32
	PlayScriptParserRULE_forControl                = 33
	PlayScriptParserRULE_forInit                   = 34
	PlayScriptParserRULE_enhancedForControl        = 35
	PlayScriptParserRULE_parExpression             = 36
	PlayScriptParserRULE_expressionList            = 37
	PlayScriptParserRULE_functionCall              = 38
	PlayScriptParserRULE_expression                = 39
	PlayScriptParserRULE_primary                   = 40
	PlayScriptParserRULE_typeList                  = 41
	PlayScriptParserRULE_typeType                  = 42
	PlayScriptParserRULE_functionType              = 43
	PlayScriptParserRULE_primitiveType             = 44
	PlayScriptParserRULE_creator                   = 45
	PlayScriptParserRULE_superSuffix               = 46
	PlayScriptParserRULE_arguments                 = 47
)

// IClassDeclarationContext is an interface to support dynamic dispatch.
type IClassDeclarationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsClassDeclarationContext differentiates from other interfaces.
	IsClassDeclarationContext()
}

type ClassDeclarationContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyClassDeclarationContext() *ClassDeclarationContext {
	var p = new(ClassDeclarationContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = PlayScriptParserRULE_classDeclaration
	return p
}

func (*ClassDeclarationContext) IsClassDeclarationContext() {}

func NewClassDeclarationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ClassDeclarationContext {
	var p = new(ClassDeclarationContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = PlayScriptParserRULE_classDeclaration

	return p
}

func (s *ClassDeclarationContext) GetParser() antlr.Parser { return s.parser }

func (s *ClassDeclarationContext) CLASS() antlr.TerminalNode {
	return s.GetToken(PlayScriptParserCLASS, 0)
}

func (s *ClassDeclarationContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(PlayScriptParserIDENTIFIER, 0)
}

func (s *ClassDeclarationContext) ClassBody() IClassBodyContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IClassBodyContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IClassBodyContext)
}

func (s *ClassDeclarationContext) EXTENDS() antlr.TerminalNode {
	return s.GetToken(PlayScriptParserEXTENDS, 0)
}

func (s *ClassDeclarationContext) TypeType() ITypeTypeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITypeTypeContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITypeTypeContext)
}

func (s *ClassDeclarationContext) IMPLEMENTS() antlr.TerminalNode {
	return s.GetToken(PlayScriptParserIMPLEMENTS, 0)
}

func (s *ClassDeclarationContext) TypeList() ITypeListContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITypeListContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITypeListContext)
}

func (s *ClassDeclarationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ClassDeclarationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ClassDeclarationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PlayScriptListener); ok {
		listenerT.EnterClassDeclaration(s)
	}
}

func (s *ClassDeclarationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PlayScriptListener); ok {
		listenerT.ExitClassDeclaration(s)
	}
}

func (s *ClassDeclarationContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PlayScriptVisitor:
		return t.VisitClassDeclaration(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *PlayScriptParser) ClassDeclaration() (localctx IClassDeclarationContext) {
	localctx = NewClassDeclarationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, PlayScriptParserRULE_classDeclaration)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(96)
		p.Match(PlayScriptParserCLASS)
	}
	{
		p.SetState(97)
		p.Match(PlayScriptParserIDENTIFIER)
	}
	p.SetState(100)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == PlayScriptParserEXTENDS {
		{
			p.SetState(98)
			p.Match(PlayScriptParserEXTENDS)
		}
		{
			p.SetState(99)
			p.TypeType()
		}

	}
	p.SetState(104)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == PlayScriptParserIMPLEMENTS {
		{
			p.SetState(102)
			p.Match(PlayScriptParserIMPLEMENTS)
		}
		{
			p.SetState(103)
			p.TypeList()
		}

	}
	{
		p.SetState(106)
		p.ClassBody()
	}

	return localctx
}

// IClassBodyContext is an interface to support dynamic dispatch.
type IClassBodyContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsClassBodyContext differentiates from other interfaces.
	IsClassBodyContext()
}

type ClassBodyContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyClassBodyContext() *ClassBodyContext {
	var p = new(ClassBodyContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = PlayScriptParserRULE_classBody
	return p
}

func (*ClassBodyContext) IsClassBodyContext() {}

func NewClassBodyContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ClassBodyContext {
	var p = new(ClassBodyContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = PlayScriptParserRULE_classBody

	return p
}

func (s *ClassBodyContext) GetParser() antlr.Parser { return s.parser }

func (s *ClassBodyContext) LBRACE() antlr.TerminalNode {
	return s.GetToken(PlayScriptParserLBRACE, 0)
}

func (s *ClassBodyContext) RBRACE() antlr.TerminalNode {
	return s.GetToken(PlayScriptParserRBRACE, 0)
}

func (s *ClassBodyContext) AllClassBodyDeclaration() []IClassBodyDeclarationContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IClassBodyDeclarationContext)(nil)).Elem())
	var tst = make([]IClassBodyDeclarationContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IClassBodyDeclarationContext)
		}
	}

	return tst
}

func (s *ClassBodyContext) ClassBodyDeclaration(i int) IClassBodyDeclarationContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IClassBodyDeclarationContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IClassBodyDeclarationContext)
}

func (s *ClassBodyContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ClassBodyContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ClassBodyContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PlayScriptListener); ok {
		listenerT.EnterClassBody(s)
	}
}

func (s *ClassBodyContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PlayScriptListener); ok {
		listenerT.ExitClassBody(s)
	}
}

func (s *ClassBodyContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PlayScriptVisitor:
		return t.VisitClassBody(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *PlayScriptParser) ClassBody() (localctx IClassBodyContext) {
	localctx = NewClassBodyContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, PlayScriptParserRULE_classBody)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(108)
		p.Match(PlayScriptParserLBRACE)
	}
	p.SetState(112)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<PlayScriptParserBOOLEAN)|(1<<PlayScriptParserBYTE)|(1<<PlayScriptParserCHAR)|(1<<PlayScriptParserCLASS)|(1<<PlayScriptParserDOUBLE)|(1<<PlayScriptParserFLOAT)|(1<<PlayScriptParserINT)|(1<<PlayScriptParserLONG))) != 0) || (((_la-37)&-(0x1f+1)) == 0 && ((1<<uint((_la-37)))&((1<<(PlayScriptParserSHORT-37))|(1<<(PlayScriptParserVOID-37))|(1<<(PlayScriptParserFUNCTION-37))|(1<<(PlayScriptParserSTRING-37)))) != 0) || _la == PlayScriptParserSEMI || _la == PlayScriptParserIDENTIFIER {
		{
			p.SetState(109)
			p.ClassBodyDeclaration()
		}

		p.SetState(114)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(115)
		p.Match(PlayScriptParserRBRACE)
	}

	return localctx
}

// IClassBodyDeclarationContext is an interface to support dynamic dispatch.
type IClassBodyDeclarationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsClassBodyDeclarationContext differentiates from other interfaces.
	IsClassBodyDeclarationContext()
}

type ClassBodyDeclarationContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyClassBodyDeclarationContext() *ClassBodyDeclarationContext {
	var p = new(ClassBodyDeclarationContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = PlayScriptParserRULE_classBodyDeclaration
	return p
}

func (*ClassBodyDeclarationContext) IsClassBodyDeclarationContext() {}

func NewClassBodyDeclarationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ClassBodyDeclarationContext {
	var p = new(ClassBodyDeclarationContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = PlayScriptParserRULE_classBodyDeclaration

	return p
}

func (s *ClassBodyDeclarationContext) GetParser() antlr.Parser { return s.parser }

func (s *ClassBodyDeclarationContext) SEMI() antlr.TerminalNode {
	return s.GetToken(PlayScriptParserSEMI, 0)
}

func (s *ClassBodyDeclarationContext) MemberDeclaration() IMemberDeclarationContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMemberDeclarationContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IMemberDeclarationContext)
}

func (s *ClassBodyDeclarationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ClassBodyDeclarationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ClassBodyDeclarationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PlayScriptListener); ok {
		listenerT.EnterClassBodyDeclaration(s)
	}
}

func (s *ClassBodyDeclarationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PlayScriptListener); ok {
		listenerT.ExitClassBodyDeclaration(s)
	}
}

func (s *ClassBodyDeclarationContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PlayScriptVisitor:
		return t.VisitClassBodyDeclaration(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *PlayScriptParser) ClassBodyDeclaration() (localctx IClassBodyDeclarationContext) {
	localctx = NewClassBodyDeclarationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, PlayScriptParserRULE_classBodyDeclaration)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(119)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case PlayScriptParserSEMI:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(117)
			p.Match(PlayScriptParserSEMI)
		}

	case PlayScriptParserBOOLEAN, PlayScriptParserBYTE, PlayScriptParserCHAR, PlayScriptParserCLASS, PlayScriptParserDOUBLE, PlayScriptParserFLOAT, PlayScriptParserINT, PlayScriptParserLONG, PlayScriptParserSHORT, PlayScriptParserVOID, PlayScriptParserFUNCTION, PlayScriptParserSTRING, PlayScriptParserIDENTIFIER:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(118)
			p.MemberDeclaration()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IMemberDeclarationContext is an interface to support dynamic dispatch.
type IMemberDeclarationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsMemberDeclarationContext differentiates from other interfaces.
	IsMemberDeclarationContext()
}

type MemberDeclarationContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMemberDeclarationContext() *MemberDeclarationContext {
	var p = new(MemberDeclarationContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = PlayScriptParserRULE_memberDeclaration
	return p
}

func (*MemberDeclarationContext) IsMemberDeclarationContext() {}

func NewMemberDeclarationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MemberDeclarationContext {
	var p = new(MemberDeclarationContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = PlayScriptParserRULE_memberDeclaration

	return p
}

func (s *MemberDeclarationContext) GetParser() antlr.Parser { return s.parser }

func (s *MemberDeclarationContext) FunctionDeclaration() IFunctionDeclarationContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFunctionDeclarationContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFunctionDeclarationContext)
}

func (s *MemberDeclarationContext) FieldDeclaration() IFieldDeclarationContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFieldDeclarationContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFieldDeclarationContext)
}

func (s *MemberDeclarationContext) ClassDeclaration() IClassDeclarationContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IClassDeclarationContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IClassDeclarationContext)
}

func (s *MemberDeclarationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MemberDeclarationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MemberDeclarationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PlayScriptListener); ok {
		listenerT.EnterMemberDeclaration(s)
	}
}

func (s *MemberDeclarationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PlayScriptListener); ok {
		listenerT.ExitMemberDeclaration(s)
	}
}

func (s *MemberDeclarationContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PlayScriptVisitor:
		return t.VisitMemberDeclaration(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *PlayScriptParser) MemberDeclaration() (localctx IMemberDeclarationContext) {
	localctx = NewMemberDeclarationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, PlayScriptParserRULE_memberDeclaration)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(124)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 4, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(121)
			p.FunctionDeclaration()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(122)
			p.FieldDeclaration()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(123)
			p.ClassDeclaration()
		}

	}

	return localctx
}

// IFunctionDeclarationContext is an interface to support dynamic dispatch.
type IFunctionDeclarationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFunctionDeclarationContext differentiates from other interfaces.
	IsFunctionDeclarationContext()
}

type FunctionDeclarationContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFunctionDeclarationContext() *FunctionDeclarationContext {
	var p = new(FunctionDeclarationContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = PlayScriptParserRULE_functionDeclaration
	return p
}

func (*FunctionDeclarationContext) IsFunctionDeclarationContext() {}

func NewFunctionDeclarationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FunctionDeclarationContext {
	var p = new(FunctionDeclarationContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = PlayScriptParserRULE_functionDeclaration

	return p
}

func (s *FunctionDeclarationContext) GetParser() antlr.Parser { return s.parser }

func (s *FunctionDeclarationContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(PlayScriptParserIDENTIFIER, 0)
}

func (s *FunctionDeclarationContext) FormalParameters() IFormalParametersContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFormalParametersContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFormalParametersContext)
}

func (s *FunctionDeclarationContext) FunctionBody() IFunctionBodyContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFunctionBodyContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFunctionBodyContext)
}

func (s *FunctionDeclarationContext) TypeTypeOrVoid() ITypeTypeOrVoidContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITypeTypeOrVoidContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITypeTypeOrVoidContext)
}

func (s *FunctionDeclarationContext) AllLBRACK() []antlr.TerminalNode {
	return s.GetTokens(PlayScriptParserLBRACK)
}

func (s *FunctionDeclarationContext) LBRACK(i int) antlr.TerminalNode {
	return s.GetToken(PlayScriptParserLBRACK, i)
}

func (s *FunctionDeclarationContext) AllRBRACK() []antlr.TerminalNode {
	return s.GetTokens(PlayScriptParserRBRACK)
}

func (s *FunctionDeclarationContext) RBRACK(i int) antlr.TerminalNode {
	return s.GetToken(PlayScriptParserRBRACK, i)
}

func (s *FunctionDeclarationContext) THROWS() antlr.TerminalNode {
	return s.GetToken(PlayScriptParserTHROWS, 0)
}

func (s *FunctionDeclarationContext) QualifiedNameList() IQualifiedNameListContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IQualifiedNameListContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IQualifiedNameListContext)
}

func (s *FunctionDeclarationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FunctionDeclarationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FunctionDeclarationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PlayScriptListener); ok {
		listenerT.EnterFunctionDeclaration(s)
	}
}

func (s *FunctionDeclarationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PlayScriptListener); ok {
		listenerT.ExitFunctionDeclaration(s)
	}
}

func (s *FunctionDeclarationContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PlayScriptVisitor:
		return t.VisitFunctionDeclaration(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *PlayScriptParser) FunctionDeclaration() (localctx IFunctionDeclarationContext) {
	localctx = NewFunctionDeclarationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, PlayScriptParserRULE_functionDeclaration)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(127)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 5, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(126)
			p.TypeTypeOrVoid()
		}

	}
	{
		p.SetState(129)
		p.Match(PlayScriptParserIDENTIFIER)
	}
	{
		p.SetState(130)
		p.FormalParameters()
	}
	p.SetState(135)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == PlayScriptParserLBRACK {
		{
			p.SetState(131)
			p.Match(PlayScriptParserLBRACK)
		}
		{
			p.SetState(132)
			p.Match(PlayScriptParserRBRACK)
		}

		p.SetState(137)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	p.SetState(140)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == PlayScriptParserTHROWS {
		{
			p.SetState(138)
			p.Match(PlayScriptParserTHROWS)
		}
		{
			p.SetState(139)
			p.QualifiedNameList()
		}

	}
	{
		p.SetState(142)
		p.FunctionBody()
	}

	return localctx
}

// IFunctionBodyContext is an interface to support dynamic dispatch.
type IFunctionBodyContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFunctionBodyContext differentiates from other interfaces.
	IsFunctionBodyContext()
}

type FunctionBodyContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFunctionBodyContext() *FunctionBodyContext {
	var p = new(FunctionBodyContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = PlayScriptParserRULE_functionBody
	return p
}

func (*FunctionBodyContext) IsFunctionBodyContext() {}

func NewFunctionBodyContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FunctionBodyContext {
	var p = new(FunctionBodyContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = PlayScriptParserRULE_functionBody

	return p
}

func (s *FunctionBodyContext) GetParser() antlr.Parser { return s.parser }

func (s *FunctionBodyContext) Block() IBlockContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBlockContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
}

func (s *FunctionBodyContext) SEMI() antlr.TerminalNode {
	return s.GetToken(PlayScriptParserSEMI, 0)
}

func (s *FunctionBodyContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FunctionBodyContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FunctionBodyContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PlayScriptListener); ok {
		listenerT.EnterFunctionBody(s)
	}
}

func (s *FunctionBodyContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PlayScriptListener); ok {
		listenerT.ExitFunctionBody(s)
	}
}

func (s *FunctionBodyContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PlayScriptVisitor:
		return t.VisitFunctionBody(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *PlayScriptParser) FunctionBody() (localctx IFunctionBodyContext) {
	localctx = NewFunctionBodyContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, PlayScriptParserRULE_functionBody)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(146)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case PlayScriptParserLBRACE:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(144)
			p.Block()
		}

	case PlayScriptParserSEMI:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(145)
			p.Match(PlayScriptParserSEMI)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// ITypeTypeOrVoidContext is an interface to support dynamic dispatch.
type ITypeTypeOrVoidContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTypeTypeOrVoidContext differentiates from other interfaces.
	IsTypeTypeOrVoidContext()
}

type TypeTypeOrVoidContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTypeTypeOrVoidContext() *TypeTypeOrVoidContext {
	var p = new(TypeTypeOrVoidContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = PlayScriptParserRULE_typeTypeOrVoid
	return p
}

func (*TypeTypeOrVoidContext) IsTypeTypeOrVoidContext() {}

func NewTypeTypeOrVoidContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TypeTypeOrVoidContext {
	var p = new(TypeTypeOrVoidContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = PlayScriptParserRULE_typeTypeOrVoid

	return p
}

func (s *TypeTypeOrVoidContext) GetParser() antlr.Parser { return s.parser }

func (s *TypeTypeOrVoidContext) TypeType() ITypeTypeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITypeTypeContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITypeTypeContext)
}

func (s *TypeTypeOrVoidContext) VOID() antlr.TerminalNode {
	return s.GetToken(PlayScriptParserVOID, 0)
}

func (s *TypeTypeOrVoidContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TypeTypeOrVoidContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TypeTypeOrVoidContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PlayScriptListener); ok {
		listenerT.EnterTypeTypeOrVoid(s)
	}
}

func (s *TypeTypeOrVoidContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PlayScriptListener); ok {
		listenerT.ExitTypeTypeOrVoid(s)
	}
}

func (s *TypeTypeOrVoidContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PlayScriptVisitor:
		return t.VisitTypeTypeOrVoid(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *PlayScriptParser) TypeTypeOrVoid() (localctx ITypeTypeOrVoidContext) {
	localctx = NewTypeTypeOrVoidContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, PlayScriptParserRULE_typeTypeOrVoid)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(150)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case PlayScriptParserBOOLEAN, PlayScriptParserBYTE, PlayScriptParserCHAR, PlayScriptParserDOUBLE, PlayScriptParserFLOAT, PlayScriptParserINT, PlayScriptParserLONG, PlayScriptParserSHORT, PlayScriptParserFUNCTION, PlayScriptParserSTRING, PlayScriptParserIDENTIFIER:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(148)
			p.TypeType()
		}

	case PlayScriptParserVOID:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(149)
			p.Match(PlayScriptParserVOID)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IQualifiedNameListContext is an interface to support dynamic dispatch.
type IQualifiedNameListContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsQualifiedNameListContext differentiates from other interfaces.
	IsQualifiedNameListContext()
}

type QualifiedNameListContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyQualifiedNameListContext() *QualifiedNameListContext {
	var p = new(QualifiedNameListContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = PlayScriptParserRULE_qualifiedNameList
	return p
}

func (*QualifiedNameListContext) IsQualifiedNameListContext() {}

func NewQualifiedNameListContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *QualifiedNameListContext {
	var p = new(QualifiedNameListContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = PlayScriptParserRULE_qualifiedNameList

	return p
}

func (s *QualifiedNameListContext) GetParser() antlr.Parser { return s.parser }

func (s *QualifiedNameListContext) AllQualifiedName() []IQualifiedNameContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IQualifiedNameContext)(nil)).Elem())
	var tst = make([]IQualifiedNameContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IQualifiedNameContext)
		}
	}

	return tst
}

func (s *QualifiedNameListContext) QualifiedName(i int) IQualifiedNameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IQualifiedNameContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IQualifiedNameContext)
}

func (s *QualifiedNameListContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(PlayScriptParserCOMMA)
}

func (s *QualifiedNameListContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(PlayScriptParserCOMMA, i)
}

func (s *QualifiedNameListContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *QualifiedNameListContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *QualifiedNameListContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PlayScriptListener); ok {
		listenerT.EnterQualifiedNameList(s)
	}
}

func (s *QualifiedNameListContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PlayScriptListener); ok {
		listenerT.ExitQualifiedNameList(s)
	}
}

func (s *QualifiedNameListContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PlayScriptVisitor:
		return t.VisitQualifiedNameList(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *PlayScriptParser) QualifiedNameList() (localctx IQualifiedNameListContext) {
	localctx = NewQualifiedNameListContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, PlayScriptParserRULE_qualifiedNameList)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(152)
		p.QualifiedName()
	}
	p.SetState(157)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == PlayScriptParserCOMMA {
		{
			p.SetState(153)
			p.Match(PlayScriptParserCOMMA)
		}
		{
			p.SetState(154)
			p.QualifiedName()
		}

		p.SetState(159)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IFormalParametersContext is an interface to support dynamic dispatch.
type IFormalParametersContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFormalParametersContext differentiates from other interfaces.
	IsFormalParametersContext()
}

type FormalParametersContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFormalParametersContext() *FormalParametersContext {
	var p = new(FormalParametersContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = PlayScriptParserRULE_formalParameters
	return p
}

func (*FormalParametersContext) IsFormalParametersContext() {}

func NewFormalParametersContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FormalParametersContext {
	var p = new(FormalParametersContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = PlayScriptParserRULE_formalParameters

	return p
}

func (s *FormalParametersContext) GetParser() antlr.Parser { return s.parser }

func (s *FormalParametersContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(PlayScriptParserLPAREN, 0)
}

func (s *FormalParametersContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(PlayScriptParserRPAREN, 0)
}

func (s *FormalParametersContext) FormalParameterList() IFormalParameterListContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFormalParameterListContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFormalParameterListContext)
}

func (s *FormalParametersContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FormalParametersContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FormalParametersContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PlayScriptListener); ok {
		listenerT.EnterFormalParameters(s)
	}
}

func (s *FormalParametersContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PlayScriptListener); ok {
		listenerT.ExitFormalParameters(s)
	}
}

func (s *FormalParametersContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PlayScriptVisitor:
		return t.VisitFormalParameters(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *PlayScriptParser) FormalParameters() (localctx IFormalParametersContext) {
	localctx = NewFormalParametersContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, PlayScriptParserRULE_formalParameters)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(160)
		p.Match(PlayScriptParserLPAREN)
	}
	p.SetState(162)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<PlayScriptParserBOOLEAN)|(1<<PlayScriptParserBYTE)|(1<<PlayScriptParserCHAR)|(1<<PlayScriptParserDOUBLE)|(1<<PlayScriptParserFINAL)|(1<<PlayScriptParserFLOAT)|(1<<PlayScriptParserINT)|(1<<PlayScriptParserLONG))) != 0) || (((_la-37)&-(0x1f+1)) == 0 && ((1<<uint((_la-37)))&((1<<(PlayScriptParserSHORT-37))|(1<<(PlayScriptParserFUNCTION-37))|(1<<(PlayScriptParserSTRING-37)))) != 0) || _la == PlayScriptParserIDENTIFIER {
		{
			p.SetState(161)
			p.FormalParameterList()
		}

	}
	{
		p.SetState(164)
		p.Match(PlayScriptParserRPAREN)
	}

	return localctx
}

// IFormalParameterListContext is an interface to support dynamic dispatch.
type IFormalParameterListContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFormalParameterListContext differentiates from other interfaces.
	IsFormalParameterListContext()
}

type FormalParameterListContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFormalParameterListContext() *FormalParameterListContext {
	var p = new(FormalParameterListContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = PlayScriptParserRULE_formalParameterList
	return p
}

func (*FormalParameterListContext) IsFormalParameterListContext() {}

func NewFormalParameterListContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FormalParameterListContext {
	var p = new(FormalParameterListContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = PlayScriptParserRULE_formalParameterList

	return p
}

func (s *FormalParameterListContext) GetParser() antlr.Parser { return s.parser }

func (s *FormalParameterListContext) AllFormalParameter() []IFormalParameterContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IFormalParameterContext)(nil)).Elem())
	var tst = make([]IFormalParameterContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IFormalParameterContext)
		}
	}

	return tst
}

func (s *FormalParameterListContext) FormalParameter(i int) IFormalParameterContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFormalParameterContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IFormalParameterContext)
}

func (s *FormalParameterListContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(PlayScriptParserCOMMA)
}

func (s *FormalParameterListContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(PlayScriptParserCOMMA, i)
}

func (s *FormalParameterListContext) LastFormalParameter() ILastFormalParameterContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ILastFormalParameterContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ILastFormalParameterContext)
}

func (s *FormalParameterListContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FormalParameterListContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FormalParameterListContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PlayScriptListener); ok {
		listenerT.EnterFormalParameterList(s)
	}
}

func (s *FormalParameterListContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PlayScriptListener); ok {
		listenerT.ExitFormalParameterList(s)
	}
}

func (s *FormalParameterListContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PlayScriptVisitor:
		return t.VisitFormalParameterList(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *PlayScriptParser) FormalParameterList() (localctx IFormalParameterListContext) {
	localctx = NewFormalParameterListContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, PlayScriptParserRULE_formalParameterList)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.SetState(179)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 14, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(166)
			p.FormalParameter()
		}
		p.SetState(171)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 12, p.GetParserRuleContext())

		for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
			if _alt == 1 {
				{
					p.SetState(167)
					p.Match(PlayScriptParserCOMMA)
				}
				{
					p.SetState(168)
					p.FormalParameter()
				}

			}
			p.SetState(173)
			p.GetErrorHandler().Sync(p)
			_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 12, p.GetParserRuleContext())
		}
		p.SetState(176)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == PlayScriptParserCOMMA {
			{
				p.SetState(174)
				p.Match(PlayScriptParserCOMMA)
			}
			{
				p.SetState(175)
				p.LastFormalParameter()
			}

		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(178)
			p.LastFormalParameter()
		}

	}

	return localctx
}

// IFormalParameterContext is an interface to support dynamic dispatch.
type IFormalParameterContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFormalParameterContext differentiates from other interfaces.
	IsFormalParameterContext()
}

type FormalParameterContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFormalParameterContext() *FormalParameterContext {
	var p = new(FormalParameterContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = PlayScriptParserRULE_formalParameter
	return p
}

func (*FormalParameterContext) IsFormalParameterContext() {}

func NewFormalParameterContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FormalParameterContext {
	var p = new(FormalParameterContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = PlayScriptParserRULE_formalParameter

	return p
}

func (s *FormalParameterContext) GetParser() antlr.Parser { return s.parser }

func (s *FormalParameterContext) TypeType() ITypeTypeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITypeTypeContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITypeTypeContext)
}

func (s *FormalParameterContext) VariableDeclaratorId() IVariableDeclaratorIdContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IVariableDeclaratorIdContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IVariableDeclaratorIdContext)
}

func (s *FormalParameterContext) AllVariableModifier() []IVariableModifierContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IVariableModifierContext)(nil)).Elem())
	var tst = make([]IVariableModifierContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IVariableModifierContext)
		}
	}

	return tst
}

func (s *FormalParameterContext) VariableModifier(i int) IVariableModifierContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IVariableModifierContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IVariableModifierContext)
}

func (s *FormalParameterContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FormalParameterContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FormalParameterContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PlayScriptListener); ok {
		listenerT.EnterFormalParameter(s)
	}
}

func (s *FormalParameterContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PlayScriptListener); ok {
		listenerT.ExitFormalParameter(s)
	}
}

func (s *FormalParameterContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PlayScriptVisitor:
		return t.VisitFormalParameter(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *PlayScriptParser) FormalParameter() (localctx IFormalParameterContext) {
	localctx = NewFormalParameterContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, PlayScriptParserRULE_formalParameter)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(184)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == PlayScriptParserFINAL {
		{
			p.SetState(181)
			p.VariableModifier()
		}

		p.SetState(186)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(187)
		p.TypeType()
	}
	{
		p.SetState(188)
		p.VariableDeclaratorId()
	}

	return localctx
}

// ILastFormalParameterContext is an interface to support dynamic dispatch.
type ILastFormalParameterContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsLastFormalParameterContext differentiates from other interfaces.
	IsLastFormalParameterContext()
}

type LastFormalParameterContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLastFormalParameterContext() *LastFormalParameterContext {
	var p = new(LastFormalParameterContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = PlayScriptParserRULE_lastFormalParameter
	return p
}

func (*LastFormalParameterContext) IsLastFormalParameterContext() {}

func NewLastFormalParameterContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LastFormalParameterContext {
	var p = new(LastFormalParameterContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = PlayScriptParserRULE_lastFormalParameter

	return p
}

func (s *LastFormalParameterContext) GetParser() antlr.Parser { return s.parser }

func (s *LastFormalParameterContext) TypeType() ITypeTypeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITypeTypeContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITypeTypeContext)
}

func (s *LastFormalParameterContext) ELLIPSIS() antlr.TerminalNode {
	return s.GetToken(PlayScriptParserELLIPSIS, 0)
}

func (s *LastFormalParameterContext) VariableDeclaratorId() IVariableDeclaratorIdContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IVariableDeclaratorIdContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IVariableDeclaratorIdContext)
}

func (s *LastFormalParameterContext) AllVariableModifier() []IVariableModifierContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IVariableModifierContext)(nil)).Elem())
	var tst = make([]IVariableModifierContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IVariableModifierContext)
		}
	}

	return tst
}

func (s *LastFormalParameterContext) VariableModifier(i int) IVariableModifierContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IVariableModifierContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IVariableModifierContext)
}

func (s *LastFormalParameterContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LastFormalParameterContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LastFormalParameterContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PlayScriptListener); ok {
		listenerT.EnterLastFormalParameter(s)
	}
}

func (s *LastFormalParameterContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PlayScriptListener); ok {
		listenerT.ExitLastFormalParameter(s)
	}
}

func (s *LastFormalParameterContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PlayScriptVisitor:
		return t.VisitLastFormalParameter(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *PlayScriptParser) LastFormalParameter() (localctx ILastFormalParameterContext) {
	localctx = NewLastFormalParameterContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, PlayScriptParserRULE_lastFormalParameter)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(193)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == PlayScriptParserFINAL {
		{
			p.SetState(190)
			p.VariableModifier()
		}

		p.SetState(195)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(196)
		p.TypeType()
	}
	{
		p.SetState(197)
		p.Match(PlayScriptParserELLIPSIS)
	}
	{
		p.SetState(198)
		p.VariableDeclaratorId()
	}

	return localctx
}

// IVariableModifierContext is an interface to support dynamic dispatch.
type IVariableModifierContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsVariableModifierContext differentiates from other interfaces.
	IsVariableModifierContext()
}

type VariableModifierContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyVariableModifierContext() *VariableModifierContext {
	var p = new(VariableModifierContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = PlayScriptParserRULE_variableModifier
	return p
}

func (*VariableModifierContext) IsVariableModifierContext() {}

func NewVariableModifierContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *VariableModifierContext {
	var p = new(VariableModifierContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = PlayScriptParserRULE_variableModifier

	return p
}

func (s *VariableModifierContext) GetParser() antlr.Parser { return s.parser }

func (s *VariableModifierContext) FINAL() antlr.TerminalNode {
	return s.GetToken(PlayScriptParserFINAL, 0)
}

func (s *VariableModifierContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VariableModifierContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *VariableModifierContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PlayScriptListener); ok {
		listenerT.EnterVariableModifier(s)
	}
}

func (s *VariableModifierContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PlayScriptListener); ok {
		listenerT.ExitVariableModifier(s)
	}
}

func (s *VariableModifierContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PlayScriptVisitor:
		return t.VisitVariableModifier(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *PlayScriptParser) VariableModifier() (localctx IVariableModifierContext) {
	localctx = NewVariableModifierContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, PlayScriptParserRULE_variableModifier)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(200)
		p.Match(PlayScriptParserFINAL)
	}

	return localctx
}

// IQualifiedNameContext is an interface to support dynamic dispatch.
type IQualifiedNameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsQualifiedNameContext differentiates from other interfaces.
	IsQualifiedNameContext()
}

type QualifiedNameContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyQualifiedNameContext() *QualifiedNameContext {
	var p = new(QualifiedNameContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = PlayScriptParserRULE_qualifiedName
	return p
}

func (*QualifiedNameContext) IsQualifiedNameContext() {}

func NewQualifiedNameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *QualifiedNameContext {
	var p = new(QualifiedNameContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = PlayScriptParserRULE_qualifiedName

	return p
}

func (s *QualifiedNameContext) GetParser() antlr.Parser { return s.parser }

func (s *QualifiedNameContext) AllIDENTIFIER() []antlr.TerminalNode {
	return s.GetTokens(PlayScriptParserIDENTIFIER)
}

func (s *QualifiedNameContext) IDENTIFIER(i int) antlr.TerminalNode {
	return s.GetToken(PlayScriptParserIDENTIFIER, i)
}

func (s *QualifiedNameContext) AllDOT() []antlr.TerminalNode {
	return s.GetTokens(PlayScriptParserDOT)
}

func (s *QualifiedNameContext) DOT(i int) antlr.TerminalNode {
	return s.GetToken(PlayScriptParserDOT, i)
}

func (s *QualifiedNameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *QualifiedNameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *QualifiedNameContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PlayScriptListener); ok {
		listenerT.EnterQualifiedName(s)
	}
}

func (s *QualifiedNameContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PlayScriptListener); ok {
		listenerT.ExitQualifiedName(s)
	}
}

func (s *QualifiedNameContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PlayScriptVisitor:
		return t.VisitQualifiedName(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *PlayScriptParser) QualifiedName() (localctx IQualifiedNameContext) {
	localctx = NewQualifiedNameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, PlayScriptParserRULE_qualifiedName)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(202)
		p.Match(PlayScriptParserIDENTIFIER)
	}
	p.SetState(207)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == PlayScriptParserDOT {
		{
			p.SetState(203)
			p.Match(PlayScriptParserDOT)
		}
		{
			p.SetState(204)
			p.Match(PlayScriptParserIDENTIFIER)
		}

		p.SetState(209)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IFieldDeclarationContext is an interface to support dynamic dispatch.
type IFieldDeclarationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFieldDeclarationContext differentiates from other interfaces.
	IsFieldDeclarationContext()
}

type FieldDeclarationContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFieldDeclarationContext() *FieldDeclarationContext {
	var p = new(FieldDeclarationContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = PlayScriptParserRULE_fieldDeclaration
	return p
}

func (*FieldDeclarationContext) IsFieldDeclarationContext() {}

func NewFieldDeclarationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FieldDeclarationContext {
	var p = new(FieldDeclarationContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = PlayScriptParserRULE_fieldDeclaration

	return p
}

func (s *FieldDeclarationContext) GetParser() antlr.Parser { return s.parser }

func (s *FieldDeclarationContext) VariableDeclarators() IVariableDeclaratorsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IVariableDeclaratorsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IVariableDeclaratorsContext)
}

func (s *FieldDeclarationContext) SEMI() antlr.TerminalNode {
	return s.GetToken(PlayScriptParserSEMI, 0)
}

func (s *FieldDeclarationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FieldDeclarationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FieldDeclarationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PlayScriptListener); ok {
		listenerT.EnterFieldDeclaration(s)
	}
}

func (s *FieldDeclarationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PlayScriptListener); ok {
		listenerT.ExitFieldDeclaration(s)
	}
}

func (s *FieldDeclarationContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PlayScriptVisitor:
		return t.VisitFieldDeclaration(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *PlayScriptParser) FieldDeclaration() (localctx IFieldDeclarationContext) {
	localctx = NewFieldDeclarationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, PlayScriptParserRULE_fieldDeclaration)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(210)
		p.VariableDeclarators()
	}
	{
		p.SetState(211)
		p.Match(PlayScriptParserSEMI)
	}

	return localctx
}

// IConstructorDeclarationContext is an interface to support dynamic dispatch.
type IConstructorDeclarationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetConstructorBody returns the constructorBody rule contexts.
	GetConstructorBody() IBlockContext

	// SetConstructorBody sets the constructorBody rule contexts.
	SetConstructorBody(IBlockContext)

	// IsConstructorDeclarationContext differentiates from other interfaces.
	IsConstructorDeclarationContext()
}

type ConstructorDeclarationContext struct {
	*antlr.BaseParserRuleContext
	parser          antlr.Parser
	constructorBody IBlockContext
}

func NewEmptyConstructorDeclarationContext() *ConstructorDeclarationContext {
	var p = new(ConstructorDeclarationContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = PlayScriptParserRULE_constructorDeclaration
	return p
}

func (*ConstructorDeclarationContext) IsConstructorDeclarationContext() {}

func NewConstructorDeclarationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ConstructorDeclarationContext {
	var p = new(ConstructorDeclarationContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = PlayScriptParserRULE_constructorDeclaration

	return p
}

func (s *ConstructorDeclarationContext) GetParser() antlr.Parser { return s.parser }

func (s *ConstructorDeclarationContext) GetConstructorBody() IBlockContext { return s.constructorBody }

func (s *ConstructorDeclarationContext) SetConstructorBody(v IBlockContext) { s.constructorBody = v }

func (s *ConstructorDeclarationContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(PlayScriptParserIDENTIFIER, 0)
}

func (s *ConstructorDeclarationContext) FormalParameters() IFormalParametersContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFormalParametersContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFormalParametersContext)
}

func (s *ConstructorDeclarationContext) Block() IBlockContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBlockContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
}

func (s *ConstructorDeclarationContext) THROWS() antlr.TerminalNode {
	return s.GetToken(PlayScriptParserTHROWS, 0)
}

func (s *ConstructorDeclarationContext) QualifiedNameList() IQualifiedNameListContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IQualifiedNameListContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IQualifiedNameListContext)
}

func (s *ConstructorDeclarationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ConstructorDeclarationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ConstructorDeclarationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PlayScriptListener); ok {
		listenerT.EnterConstructorDeclaration(s)
	}
}

func (s *ConstructorDeclarationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PlayScriptListener); ok {
		listenerT.ExitConstructorDeclaration(s)
	}
}

func (s *ConstructorDeclarationContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PlayScriptVisitor:
		return t.VisitConstructorDeclaration(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *PlayScriptParser) ConstructorDeclaration() (localctx IConstructorDeclarationContext) {
	localctx = NewConstructorDeclarationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 30, PlayScriptParserRULE_constructorDeclaration)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(213)
		p.Match(PlayScriptParserIDENTIFIER)
	}
	{
		p.SetState(214)
		p.FormalParameters()
	}
	p.SetState(217)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == PlayScriptParserTHROWS {
		{
			p.SetState(215)
			p.Match(PlayScriptParserTHROWS)
		}
		{
			p.SetState(216)
			p.QualifiedNameList()
		}

	}
	{
		p.SetState(219)

		var _x = p.Block()

		localctx.(*ConstructorDeclarationContext).constructorBody = _x
	}

	return localctx
}

// IVariableDeclaratorsContext is an interface to support dynamic dispatch.
type IVariableDeclaratorsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsVariableDeclaratorsContext differentiates from other interfaces.
	IsVariableDeclaratorsContext()
}

type VariableDeclaratorsContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyVariableDeclaratorsContext() *VariableDeclaratorsContext {
	var p = new(VariableDeclaratorsContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = PlayScriptParserRULE_variableDeclarators
	return p
}

func (*VariableDeclaratorsContext) IsVariableDeclaratorsContext() {}

func NewVariableDeclaratorsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *VariableDeclaratorsContext {
	var p = new(VariableDeclaratorsContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = PlayScriptParserRULE_variableDeclarators

	return p
}

func (s *VariableDeclaratorsContext) GetParser() antlr.Parser { return s.parser }

func (s *VariableDeclaratorsContext) TypeType() ITypeTypeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITypeTypeContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITypeTypeContext)
}

func (s *VariableDeclaratorsContext) AllVariableDeclarator() []IVariableDeclaratorContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IVariableDeclaratorContext)(nil)).Elem())
	var tst = make([]IVariableDeclaratorContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IVariableDeclaratorContext)
		}
	}

	return tst
}

func (s *VariableDeclaratorsContext) VariableDeclarator(i int) IVariableDeclaratorContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IVariableDeclaratorContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IVariableDeclaratorContext)
}

func (s *VariableDeclaratorsContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(PlayScriptParserCOMMA)
}

func (s *VariableDeclaratorsContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(PlayScriptParserCOMMA, i)
}

func (s *VariableDeclaratorsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VariableDeclaratorsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *VariableDeclaratorsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PlayScriptListener); ok {
		listenerT.EnterVariableDeclarators(s)
	}
}

func (s *VariableDeclaratorsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PlayScriptListener); ok {
		listenerT.ExitVariableDeclarators(s)
	}
}

func (s *VariableDeclaratorsContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PlayScriptVisitor:
		return t.VisitVariableDeclarators(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *PlayScriptParser) VariableDeclarators() (localctx IVariableDeclaratorsContext) {
	localctx = NewVariableDeclaratorsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 32, PlayScriptParserRULE_variableDeclarators)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(221)
		p.TypeType()
	}
	{
		p.SetState(222)
		p.VariableDeclarator()
	}
	p.SetState(227)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == PlayScriptParserCOMMA {
		{
			p.SetState(223)
			p.Match(PlayScriptParserCOMMA)
		}
		{
			p.SetState(224)
			p.VariableDeclarator()
		}

		p.SetState(229)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IVariableDeclaratorContext is an interface to support dynamic dispatch.
type IVariableDeclaratorContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsVariableDeclaratorContext differentiates from other interfaces.
	IsVariableDeclaratorContext()
}

type VariableDeclaratorContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyVariableDeclaratorContext() *VariableDeclaratorContext {
	var p = new(VariableDeclaratorContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = PlayScriptParserRULE_variableDeclarator
	return p
}

func (*VariableDeclaratorContext) IsVariableDeclaratorContext() {}

func NewVariableDeclaratorContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *VariableDeclaratorContext {
	var p = new(VariableDeclaratorContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = PlayScriptParserRULE_variableDeclarator

	return p
}

func (s *VariableDeclaratorContext) GetParser() antlr.Parser { return s.parser }

func (s *VariableDeclaratorContext) VariableDeclaratorId() IVariableDeclaratorIdContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IVariableDeclaratorIdContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IVariableDeclaratorIdContext)
}

func (s *VariableDeclaratorContext) ASSIGN() antlr.TerminalNode {
	return s.GetToken(PlayScriptParserASSIGN, 0)
}

func (s *VariableDeclaratorContext) VariableInitializer() IVariableInitializerContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IVariableInitializerContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IVariableInitializerContext)
}

func (s *VariableDeclaratorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VariableDeclaratorContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *VariableDeclaratorContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PlayScriptListener); ok {
		listenerT.EnterVariableDeclarator(s)
	}
}

func (s *VariableDeclaratorContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PlayScriptListener); ok {
		listenerT.ExitVariableDeclarator(s)
	}
}

func (s *VariableDeclaratorContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PlayScriptVisitor:
		return t.VisitVariableDeclarator(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *PlayScriptParser) VariableDeclarator() (localctx IVariableDeclaratorContext) {
	localctx = NewVariableDeclaratorContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 34, PlayScriptParserRULE_variableDeclarator)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(230)
		p.VariableDeclaratorId()
	}
	p.SetState(233)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == PlayScriptParserASSIGN {
		{
			p.SetState(231)
			p.Match(PlayScriptParserASSIGN)
		}
		{
			p.SetState(232)
			p.VariableInitializer()
		}

	}

	return localctx
}

// IVariableDeclaratorIdContext is an interface to support dynamic dispatch.
type IVariableDeclaratorIdContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsVariableDeclaratorIdContext differentiates from other interfaces.
	IsVariableDeclaratorIdContext()
}

type VariableDeclaratorIdContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyVariableDeclaratorIdContext() *VariableDeclaratorIdContext {
	var p = new(VariableDeclaratorIdContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = PlayScriptParserRULE_variableDeclaratorId
	return p
}

func (*VariableDeclaratorIdContext) IsVariableDeclaratorIdContext() {}

func NewVariableDeclaratorIdContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *VariableDeclaratorIdContext {
	var p = new(VariableDeclaratorIdContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = PlayScriptParserRULE_variableDeclaratorId

	return p
}

func (s *VariableDeclaratorIdContext) GetParser() antlr.Parser { return s.parser }

func (s *VariableDeclaratorIdContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(PlayScriptParserIDENTIFIER, 0)
}

func (s *VariableDeclaratorIdContext) AllLBRACK() []antlr.TerminalNode {
	return s.GetTokens(PlayScriptParserLBRACK)
}

func (s *VariableDeclaratorIdContext) LBRACK(i int) antlr.TerminalNode {
	return s.GetToken(PlayScriptParserLBRACK, i)
}

func (s *VariableDeclaratorIdContext) AllRBRACK() []antlr.TerminalNode {
	return s.GetTokens(PlayScriptParserRBRACK)
}

func (s *VariableDeclaratorIdContext) RBRACK(i int) antlr.TerminalNode {
	return s.GetToken(PlayScriptParserRBRACK, i)
}

func (s *VariableDeclaratorIdContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VariableDeclaratorIdContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *VariableDeclaratorIdContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PlayScriptListener); ok {
		listenerT.EnterVariableDeclaratorId(s)
	}
}

func (s *VariableDeclaratorIdContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PlayScriptListener); ok {
		listenerT.ExitVariableDeclaratorId(s)
	}
}

func (s *VariableDeclaratorIdContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PlayScriptVisitor:
		return t.VisitVariableDeclaratorId(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *PlayScriptParser) VariableDeclaratorId() (localctx IVariableDeclaratorIdContext) {
	localctx = NewVariableDeclaratorIdContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 36, PlayScriptParserRULE_variableDeclaratorId)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(235)
		p.Match(PlayScriptParserIDENTIFIER)
	}
	p.SetState(240)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == PlayScriptParserLBRACK {
		{
			p.SetState(236)
			p.Match(PlayScriptParserLBRACK)
		}
		{
			p.SetState(237)
			p.Match(PlayScriptParserRBRACK)
		}

		p.SetState(242)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IVariableInitializerContext is an interface to support dynamic dispatch.
type IVariableInitializerContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsVariableInitializerContext differentiates from other interfaces.
	IsVariableInitializerContext()
}

type VariableInitializerContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyVariableInitializerContext() *VariableInitializerContext {
	var p = new(VariableInitializerContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = PlayScriptParserRULE_variableInitializer
	return p
}

func (*VariableInitializerContext) IsVariableInitializerContext() {}

func NewVariableInitializerContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *VariableInitializerContext {
	var p = new(VariableInitializerContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = PlayScriptParserRULE_variableInitializer

	return p
}

func (s *VariableInitializerContext) GetParser() antlr.Parser { return s.parser }

func (s *VariableInitializerContext) ArrayInitializer() IArrayInitializerContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IArrayInitializerContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IArrayInitializerContext)
}

func (s *VariableInitializerContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *VariableInitializerContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VariableInitializerContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *VariableInitializerContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PlayScriptListener); ok {
		listenerT.EnterVariableInitializer(s)
	}
}

func (s *VariableInitializerContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PlayScriptListener); ok {
		listenerT.ExitVariableInitializer(s)
	}
}

func (s *VariableInitializerContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PlayScriptVisitor:
		return t.VisitVariableInitializer(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *PlayScriptParser) VariableInitializer() (localctx IVariableInitializerContext) {
	localctx = NewVariableInitializerContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 38, PlayScriptParserRULE_variableInitializer)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(245)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case PlayScriptParserLBRACE:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(243)
			p.ArrayInitializer()
		}

	case PlayScriptParserSUPER, PlayScriptParserTHIS, PlayScriptParserDECIMAL_LITERAL, PlayScriptParserHEX_LITERAL, PlayScriptParserOCT_LITERAL, PlayScriptParserBINARY_LITERAL, PlayScriptParserFLOAT_LITERAL, PlayScriptParserHEX_FLOAT_LITERAL, PlayScriptParserBOOL_LITERAL, PlayScriptParserCHAR_LITERAL, PlayScriptParserSTRING_LITERAL, PlayScriptParserNULL_LITERAL, PlayScriptParserLPAREN, PlayScriptParserBANG, PlayScriptParserTILDE, PlayScriptParserINC, PlayScriptParserDEC, PlayScriptParserADD, PlayScriptParserSUB, PlayScriptParserIDENTIFIER:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(244)
			p.expression(0)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IArrayInitializerContext is an interface to support dynamic dispatch.
type IArrayInitializerContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsArrayInitializerContext differentiates from other interfaces.
	IsArrayInitializerContext()
}

type ArrayInitializerContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyArrayInitializerContext() *ArrayInitializerContext {
	var p = new(ArrayInitializerContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = PlayScriptParserRULE_arrayInitializer
	return p
}

func (*ArrayInitializerContext) IsArrayInitializerContext() {}

func NewArrayInitializerContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ArrayInitializerContext {
	var p = new(ArrayInitializerContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = PlayScriptParserRULE_arrayInitializer

	return p
}

func (s *ArrayInitializerContext) GetParser() antlr.Parser { return s.parser }

func (s *ArrayInitializerContext) LBRACE() antlr.TerminalNode {
	return s.GetToken(PlayScriptParserLBRACE, 0)
}

func (s *ArrayInitializerContext) RBRACE() antlr.TerminalNode {
	return s.GetToken(PlayScriptParserRBRACE, 0)
}

func (s *ArrayInitializerContext) AllVariableInitializer() []IVariableInitializerContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IVariableInitializerContext)(nil)).Elem())
	var tst = make([]IVariableInitializerContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IVariableInitializerContext)
		}
	}

	return tst
}

func (s *ArrayInitializerContext) VariableInitializer(i int) IVariableInitializerContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IVariableInitializerContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IVariableInitializerContext)
}

func (s *ArrayInitializerContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(PlayScriptParserCOMMA)
}

func (s *ArrayInitializerContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(PlayScriptParserCOMMA, i)
}

func (s *ArrayInitializerContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ArrayInitializerContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ArrayInitializerContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PlayScriptListener); ok {
		listenerT.EnterArrayInitializer(s)
	}
}

func (s *ArrayInitializerContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PlayScriptListener); ok {
		listenerT.ExitArrayInitializer(s)
	}
}

func (s *ArrayInitializerContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PlayScriptVisitor:
		return t.VisitArrayInitializer(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *PlayScriptParser) ArrayInitializer() (localctx IArrayInitializerContext) {
	localctx = NewArrayInitializerContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 40, PlayScriptParserRULE_arrayInitializer)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(247)
		p.Match(PlayScriptParserLBRACE)
	}
	p.SetState(259)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if (((_la-40)&-(0x1f+1)) == 0 && ((1<<uint((_la-40)))&((1<<(PlayScriptParserSUPER-40))|(1<<(PlayScriptParserTHIS-40))|(1<<(PlayScriptParserDECIMAL_LITERAL-40))|(1<<(PlayScriptParserHEX_LITERAL-40))|(1<<(PlayScriptParserOCT_LITERAL-40))|(1<<(PlayScriptParserBINARY_LITERAL-40))|(1<<(PlayScriptParserFLOAT_LITERAL-40))|(1<<(PlayScriptParserHEX_FLOAT_LITERAL-40))|(1<<(PlayScriptParserBOOL_LITERAL-40))|(1<<(PlayScriptParserCHAR_LITERAL-40))|(1<<(PlayScriptParserSTRING_LITERAL-40))|(1<<(PlayScriptParserNULL_LITERAL-40))|(1<<(PlayScriptParserLPAREN-40))|(1<<(PlayScriptParserLBRACE-40)))) != 0) || (((_la-75)&-(0x1f+1)) == 0 && ((1<<uint((_la-75)))&((1<<(PlayScriptParserBANG-75))|(1<<(PlayScriptParserTILDE-75))|(1<<(PlayScriptParserINC-75))|(1<<(PlayScriptParserDEC-75))|(1<<(PlayScriptParserADD-75))|(1<<(PlayScriptParserSUB-75)))) != 0) || _la == PlayScriptParserIDENTIFIER {
		{
			p.SetState(248)
			p.VariableInitializer()
		}
		p.SetState(253)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 23, p.GetParserRuleContext())

		for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
			if _alt == 1 {
				{
					p.SetState(249)
					p.Match(PlayScriptParserCOMMA)
				}
				{
					p.SetState(250)
					p.VariableInitializer()
				}

			}
			p.SetState(255)
			p.GetErrorHandler().Sync(p)
			_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 23, p.GetParserRuleContext())
		}
		p.SetState(257)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == PlayScriptParserCOMMA {
			{
				p.SetState(256)
				p.Match(PlayScriptParserCOMMA)
			}

		}

	}
	{
		p.SetState(261)
		p.Match(PlayScriptParserRBRACE)
	}

	return localctx
}

// IClassOrInterfaceTypeContext is an interface to support dynamic dispatch.
type IClassOrInterfaceTypeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsClassOrInterfaceTypeContext differentiates from other interfaces.
	IsClassOrInterfaceTypeContext()
}

type ClassOrInterfaceTypeContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyClassOrInterfaceTypeContext() *ClassOrInterfaceTypeContext {
	var p = new(ClassOrInterfaceTypeContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = PlayScriptParserRULE_classOrInterfaceType
	return p
}

func (*ClassOrInterfaceTypeContext) IsClassOrInterfaceTypeContext() {}

func NewClassOrInterfaceTypeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ClassOrInterfaceTypeContext {
	var p = new(ClassOrInterfaceTypeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = PlayScriptParserRULE_classOrInterfaceType

	return p
}

func (s *ClassOrInterfaceTypeContext) GetParser() antlr.Parser { return s.parser }

func (s *ClassOrInterfaceTypeContext) AllIDENTIFIER() []antlr.TerminalNode {
	return s.GetTokens(PlayScriptParserIDENTIFIER)
}

func (s *ClassOrInterfaceTypeContext) IDENTIFIER(i int) antlr.TerminalNode {
	return s.GetToken(PlayScriptParserIDENTIFIER, i)
}

func (s *ClassOrInterfaceTypeContext) AllDOT() []antlr.TerminalNode {
	return s.GetTokens(PlayScriptParserDOT)
}

func (s *ClassOrInterfaceTypeContext) DOT(i int) antlr.TerminalNode {
	return s.GetToken(PlayScriptParserDOT, i)
}

func (s *ClassOrInterfaceTypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ClassOrInterfaceTypeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ClassOrInterfaceTypeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PlayScriptListener); ok {
		listenerT.EnterClassOrInterfaceType(s)
	}
}

func (s *ClassOrInterfaceTypeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PlayScriptListener); ok {
		listenerT.ExitClassOrInterfaceType(s)
	}
}

func (s *ClassOrInterfaceTypeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PlayScriptVisitor:
		return t.VisitClassOrInterfaceType(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *PlayScriptParser) ClassOrInterfaceType() (localctx IClassOrInterfaceTypeContext) {
	localctx = NewClassOrInterfaceTypeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 42, PlayScriptParserRULE_classOrInterfaceType)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(263)
		p.Match(PlayScriptParserIDENTIFIER)
	}
	p.SetState(268)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 26, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(264)
				p.Match(PlayScriptParserDOT)
			}
			{
				p.SetState(265)
				p.Match(PlayScriptParserIDENTIFIER)
			}

		}
		p.SetState(270)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 26, p.GetParserRuleContext())
	}

	return localctx
}

// ITypeArgumentContext is an interface to support dynamic dispatch.
type ITypeArgumentContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTypeArgumentContext differentiates from other interfaces.
	IsTypeArgumentContext()
}

type TypeArgumentContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTypeArgumentContext() *TypeArgumentContext {
	var p = new(TypeArgumentContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = PlayScriptParserRULE_typeArgument
	return p
}

func (*TypeArgumentContext) IsTypeArgumentContext() {}

func NewTypeArgumentContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TypeArgumentContext {
	var p = new(TypeArgumentContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = PlayScriptParserRULE_typeArgument

	return p
}

func (s *TypeArgumentContext) GetParser() antlr.Parser { return s.parser }

func (s *TypeArgumentContext) TypeType() ITypeTypeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITypeTypeContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITypeTypeContext)
}

func (s *TypeArgumentContext) QUESTION() antlr.TerminalNode {
	return s.GetToken(PlayScriptParserQUESTION, 0)
}

func (s *TypeArgumentContext) EXTENDS() antlr.TerminalNode {
	return s.GetToken(PlayScriptParserEXTENDS, 0)
}

func (s *TypeArgumentContext) SUPER() antlr.TerminalNode {
	return s.GetToken(PlayScriptParserSUPER, 0)
}

func (s *TypeArgumentContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TypeArgumentContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TypeArgumentContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PlayScriptListener); ok {
		listenerT.EnterTypeArgument(s)
	}
}

func (s *TypeArgumentContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PlayScriptListener); ok {
		listenerT.ExitTypeArgument(s)
	}
}

func (s *TypeArgumentContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PlayScriptVisitor:
		return t.VisitTypeArgument(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *PlayScriptParser) TypeArgument() (localctx ITypeArgumentContext) {
	localctx = NewTypeArgumentContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 44, PlayScriptParserRULE_typeArgument)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(277)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case PlayScriptParserBOOLEAN, PlayScriptParserBYTE, PlayScriptParserCHAR, PlayScriptParserDOUBLE, PlayScriptParserFLOAT, PlayScriptParserINT, PlayScriptParserLONG, PlayScriptParserSHORT, PlayScriptParserFUNCTION, PlayScriptParserSTRING, PlayScriptParserIDENTIFIER:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(271)
			p.TypeType()
		}

	case PlayScriptParserQUESTION:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(272)
			p.Match(PlayScriptParserQUESTION)
		}
		p.SetState(275)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == PlayScriptParserEXTENDS || _la == PlayScriptParserSUPER {
			{
				p.SetState(273)
				_la = p.GetTokenStream().LA(1)

				if !(_la == PlayScriptParserEXTENDS || _la == PlayScriptParserSUPER) {
					p.GetErrorHandler().RecoverInline(p)
				} else {
					p.GetErrorHandler().ReportMatch(p)
					p.Consume()
				}
			}
			{
				p.SetState(274)
				p.TypeType()
			}

		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// ILiteralContext is an interface to support dynamic dispatch.
type ILiteralContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsLiteralContext differentiates from other interfaces.
	IsLiteralContext()
}

type LiteralContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLiteralContext() *LiteralContext {
	var p = new(LiteralContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = PlayScriptParserRULE_literal
	return p
}

func (*LiteralContext) IsLiteralContext() {}

func NewLiteralContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LiteralContext {
	var p = new(LiteralContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = PlayScriptParserRULE_literal

	return p
}

func (s *LiteralContext) GetParser() antlr.Parser { return s.parser }

func (s *LiteralContext) IntegerLiteral() IIntegerLiteralContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIntegerLiteralContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIntegerLiteralContext)
}

func (s *LiteralContext) FloatLiteral() IFloatLiteralContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFloatLiteralContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFloatLiteralContext)
}

func (s *LiteralContext) CHAR_LITERAL() antlr.TerminalNode {
	return s.GetToken(PlayScriptParserCHAR_LITERAL, 0)
}

func (s *LiteralContext) STRING_LITERAL() antlr.TerminalNode {
	return s.GetToken(PlayScriptParserSTRING_LITERAL, 0)
}

func (s *LiteralContext) BOOL_LITERAL() antlr.TerminalNode {
	return s.GetToken(PlayScriptParserBOOL_LITERAL, 0)
}

func (s *LiteralContext) NULL_LITERAL() antlr.TerminalNode {
	return s.GetToken(PlayScriptParserNULL_LITERAL, 0)
}

func (s *LiteralContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LiteralContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LiteralContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PlayScriptListener); ok {
		listenerT.EnterLiteral(s)
	}
}

func (s *LiteralContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PlayScriptListener); ok {
		listenerT.ExitLiteral(s)
	}
}

func (s *LiteralContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PlayScriptVisitor:
		return t.VisitLiteral(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *PlayScriptParser) Literal() (localctx ILiteralContext) {
	localctx = NewLiteralContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 46, PlayScriptParserRULE_literal)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(285)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case PlayScriptParserDECIMAL_LITERAL, PlayScriptParserHEX_LITERAL, PlayScriptParserOCT_LITERAL, PlayScriptParserBINARY_LITERAL:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(279)
			p.IntegerLiteral()
		}

	case PlayScriptParserFLOAT_LITERAL, PlayScriptParserHEX_FLOAT_LITERAL:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(280)
			p.FloatLiteral()
		}

	case PlayScriptParserCHAR_LITERAL:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(281)
			p.Match(PlayScriptParserCHAR_LITERAL)
		}

	case PlayScriptParserSTRING_LITERAL:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(282)
			p.Match(PlayScriptParserSTRING_LITERAL)
		}

	case PlayScriptParserBOOL_LITERAL:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(283)
			p.Match(PlayScriptParserBOOL_LITERAL)
		}

	case PlayScriptParserNULL_LITERAL:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(284)
			p.Match(PlayScriptParserNULL_LITERAL)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IIntegerLiteralContext is an interface to support dynamic dispatch.
type IIntegerLiteralContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsIntegerLiteralContext differentiates from other interfaces.
	IsIntegerLiteralContext()
}

type IntegerLiteralContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIntegerLiteralContext() *IntegerLiteralContext {
	var p = new(IntegerLiteralContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = PlayScriptParserRULE_integerLiteral
	return p
}

func (*IntegerLiteralContext) IsIntegerLiteralContext() {}

func NewIntegerLiteralContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IntegerLiteralContext {
	var p = new(IntegerLiteralContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = PlayScriptParserRULE_integerLiteral

	return p
}

func (s *IntegerLiteralContext) GetParser() antlr.Parser { return s.parser }

func (s *IntegerLiteralContext) DECIMAL_LITERAL() antlr.TerminalNode {
	return s.GetToken(PlayScriptParserDECIMAL_LITERAL, 0)
}

func (s *IntegerLiteralContext) HEX_LITERAL() antlr.TerminalNode {
	return s.GetToken(PlayScriptParserHEX_LITERAL, 0)
}

func (s *IntegerLiteralContext) OCT_LITERAL() antlr.TerminalNode {
	return s.GetToken(PlayScriptParserOCT_LITERAL, 0)
}

func (s *IntegerLiteralContext) BINARY_LITERAL() antlr.TerminalNode {
	return s.GetToken(PlayScriptParserBINARY_LITERAL, 0)
}

func (s *IntegerLiteralContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IntegerLiteralContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IntegerLiteralContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PlayScriptListener); ok {
		listenerT.EnterIntegerLiteral(s)
	}
}

func (s *IntegerLiteralContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PlayScriptListener); ok {
		listenerT.ExitIntegerLiteral(s)
	}
}

func (s *IntegerLiteralContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PlayScriptVisitor:
		return t.VisitIntegerLiteral(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *PlayScriptParser) IntegerLiteral() (localctx IIntegerLiteralContext) {
	localctx = NewIntegerLiteralContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 48, PlayScriptParserRULE_integerLiteral)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(287)
		_la = p.GetTokenStream().LA(1)

		if !(((_la-53)&-(0x1f+1)) == 0 && ((1<<uint((_la-53)))&((1<<(PlayScriptParserDECIMAL_LITERAL-53))|(1<<(PlayScriptParserHEX_LITERAL-53))|(1<<(PlayScriptParserOCT_LITERAL-53))|(1<<(PlayScriptParserBINARY_LITERAL-53)))) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// IFloatLiteralContext is an interface to support dynamic dispatch.
type IFloatLiteralContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFloatLiteralContext differentiates from other interfaces.
	IsFloatLiteralContext()
}

type FloatLiteralContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFloatLiteralContext() *FloatLiteralContext {
	var p = new(FloatLiteralContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = PlayScriptParserRULE_floatLiteral
	return p
}

func (*FloatLiteralContext) IsFloatLiteralContext() {}

func NewFloatLiteralContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FloatLiteralContext {
	var p = new(FloatLiteralContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = PlayScriptParserRULE_floatLiteral

	return p
}

func (s *FloatLiteralContext) GetParser() antlr.Parser { return s.parser }

func (s *FloatLiteralContext) FLOAT_LITERAL() antlr.TerminalNode {
	return s.GetToken(PlayScriptParserFLOAT_LITERAL, 0)
}

func (s *FloatLiteralContext) HEX_FLOAT_LITERAL() antlr.TerminalNode {
	return s.GetToken(PlayScriptParserHEX_FLOAT_LITERAL, 0)
}

func (s *FloatLiteralContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FloatLiteralContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FloatLiteralContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PlayScriptListener); ok {
		listenerT.EnterFloatLiteral(s)
	}
}

func (s *FloatLiteralContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PlayScriptListener); ok {
		listenerT.ExitFloatLiteral(s)
	}
}

func (s *FloatLiteralContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PlayScriptVisitor:
		return t.VisitFloatLiteral(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *PlayScriptParser) FloatLiteral() (localctx IFloatLiteralContext) {
	localctx = NewFloatLiteralContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 50, PlayScriptParserRULE_floatLiteral)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(289)
		_la = p.GetTokenStream().LA(1)

		if !(_la == PlayScriptParserFLOAT_LITERAL || _la == PlayScriptParserHEX_FLOAT_LITERAL) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// IProgContext is an interface to support dynamic dispatch.
type IProgContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsProgContext differentiates from other interfaces.
	IsProgContext()
}

type ProgContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyProgContext() *ProgContext {
	var p = new(ProgContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = PlayScriptParserRULE_prog
	return p
}

func (*ProgContext) IsProgContext() {}

func NewProgContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ProgContext {
	var p = new(ProgContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = PlayScriptParserRULE_prog

	return p
}

func (s *ProgContext) GetParser() antlr.Parser { return s.parser }

func (s *ProgContext) BlockStatements() IBlockStatementsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBlockStatementsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBlockStatementsContext)
}

func (s *ProgContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ProgContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ProgContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PlayScriptListener); ok {
		listenerT.EnterProg(s)
	}
}

func (s *ProgContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PlayScriptListener); ok {
		listenerT.ExitProg(s)
	}
}

func (s *ProgContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PlayScriptVisitor:
		return t.VisitProg(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *PlayScriptParser) Prog() (localctx IProgContext) {
	localctx = NewProgContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 52, PlayScriptParserRULE_prog)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(291)
		p.BlockStatements()
	}

	return localctx
}

// IBlockContext is an interface to support dynamic dispatch.
type IBlockContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsBlockContext differentiates from other interfaces.
	IsBlockContext()
}

type BlockContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBlockContext() *BlockContext {
	var p = new(BlockContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = PlayScriptParserRULE_block
	return p
}

func (*BlockContext) IsBlockContext() {}

func NewBlockContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BlockContext {
	var p = new(BlockContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = PlayScriptParserRULE_block

	return p
}

func (s *BlockContext) GetParser() antlr.Parser { return s.parser }

func (s *BlockContext) LBRACE() antlr.TerminalNode {
	return s.GetToken(PlayScriptParserLBRACE, 0)
}

func (s *BlockContext) BlockStatements() IBlockStatementsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBlockStatementsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBlockStatementsContext)
}

func (s *BlockContext) RBRACE() antlr.TerminalNode {
	return s.GetToken(PlayScriptParserRBRACE, 0)
}

func (s *BlockContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BlockContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BlockContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PlayScriptListener); ok {
		listenerT.EnterBlock(s)
	}
}

func (s *BlockContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PlayScriptListener); ok {
		listenerT.ExitBlock(s)
	}
}

func (s *BlockContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PlayScriptVisitor:
		return t.VisitBlock(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *PlayScriptParser) Block() (localctx IBlockContext) {
	localctx = NewBlockContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 54, PlayScriptParserRULE_block)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(293)
		p.Match(PlayScriptParserLBRACE)
	}
	{
		p.SetState(294)
		p.BlockStatements()
	}
	{
		p.SetState(295)
		p.Match(PlayScriptParserRBRACE)
	}

	return localctx
}

// IBlockStatementsContext is an interface to support dynamic dispatch.
type IBlockStatementsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsBlockStatementsContext differentiates from other interfaces.
	IsBlockStatementsContext()
}

type BlockStatementsContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBlockStatementsContext() *BlockStatementsContext {
	var p = new(BlockStatementsContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = PlayScriptParserRULE_blockStatements
	return p
}

func (*BlockStatementsContext) IsBlockStatementsContext() {}

func NewBlockStatementsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BlockStatementsContext {
	var p = new(BlockStatementsContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = PlayScriptParserRULE_blockStatements

	return p
}

func (s *BlockStatementsContext) GetParser() antlr.Parser { return s.parser }

func (s *BlockStatementsContext) AllBlockStatement() []IBlockStatementContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IBlockStatementContext)(nil)).Elem())
	var tst = make([]IBlockStatementContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IBlockStatementContext)
		}
	}

	return tst
}

func (s *BlockStatementsContext) BlockStatement(i int) IBlockStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBlockStatementContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IBlockStatementContext)
}

func (s *BlockStatementsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BlockStatementsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BlockStatementsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PlayScriptListener); ok {
		listenerT.EnterBlockStatements(s)
	}
}

func (s *BlockStatementsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PlayScriptListener); ok {
		listenerT.ExitBlockStatements(s)
	}
}

func (s *BlockStatementsContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PlayScriptVisitor:
		return t.VisitBlockStatements(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *PlayScriptParser) BlockStatements() (localctx IBlockStatementsContext) {
	localctx = NewBlockStatementsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 56, PlayScriptParserRULE_blockStatements)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(300)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<PlayScriptParserBOOLEAN)|(1<<PlayScriptParserBREAK)|(1<<PlayScriptParserBYTE)|(1<<PlayScriptParserCHAR)|(1<<PlayScriptParserCLASS)|(1<<PlayScriptParserCONTINUE)|(1<<PlayScriptParserDO)|(1<<PlayScriptParserDOUBLE)|(1<<PlayScriptParserFLOAT)|(1<<PlayScriptParserFOR)|(1<<PlayScriptParserIF)|(1<<PlayScriptParserINT)|(1<<PlayScriptParserLONG))) != 0) || (((_la-36)&-(0x1f+1)) == 0 && ((1<<uint((_la-36)))&((1<<(PlayScriptParserRETURN-36))|(1<<(PlayScriptParserSHORT-36))|(1<<(PlayScriptParserSUPER-36))|(1<<(PlayScriptParserSWITCH-36))|(1<<(PlayScriptParserTHIS-36))|(1<<(PlayScriptParserVOID-36))|(1<<(PlayScriptParserWHILE-36))|(1<<(PlayScriptParserFUNCTION-36))|(1<<(PlayScriptParserSTRING-36))|(1<<(PlayScriptParserDECIMAL_LITERAL-36))|(1<<(PlayScriptParserHEX_LITERAL-36))|(1<<(PlayScriptParserOCT_LITERAL-36))|(1<<(PlayScriptParserBINARY_LITERAL-36))|(1<<(PlayScriptParserFLOAT_LITERAL-36))|(1<<(PlayScriptParserHEX_FLOAT_LITERAL-36))|(1<<(PlayScriptParserBOOL_LITERAL-36))|(1<<(PlayScriptParserCHAR_LITERAL-36))|(1<<(PlayScriptParserSTRING_LITERAL-36))|(1<<(PlayScriptParserNULL_LITERAL-36))|(1<<(PlayScriptParserLPAREN-36))|(1<<(PlayScriptParserLBRACE-36)))) != 0) || (((_la-69)&-(0x1f+1)) == 0 && ((1<<uint((_la-69)))&((1<<(PlayScriptParserSEMI-69))|(1<<(PlayScriptParserBANG-69))|(1<<(PlayScriptParserTILDE-69))|(1<<(PlayScriptParserINC-69))|(1<<(PlayScriptParserDEC-69))|(1<<(PlayScriptParserADD-69))|(1<<(PlayScriptParserSUB-69)))) != 0) || _la == PlayScriptParserIDENTIFIER {
		{
			p.SetState(297)
			p.BlockStatement()
		}

		p.SetState(302)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IBlockStatementContext is an interface to support dynamic dispatch.
type IBlockStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsBlockStatementContext differentiates from other interfaces.
	IsBlockStatementContext()
}

type BlockStatementContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBlockStatementContext() *BlockStatementContext {
	var p = new(BlockStatementContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = PlayScriptParserRULE_blockStatement
	return p
}

func (*BlockStatementContext) IsBlockStatementContext() {}

func NewBlockStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BlockStatementContext {
	var p = new(BlockStatementContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = PlayScriptParserRULE_blockStatement

	return p
}

func (s *BlockStatementContext) GetParser() antlr.Parser { return s.parser }

func (s *BlockStatementContext) VariableDeclarators() IVariableDeclaratorsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IVariableDeclaratorsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IVariableDeclaratorsContext)
}

func (s *BlockStatementContext) SEMI() antlr.TerminalNode {
	return s.GetToken(PlayScriptParserSEMI, 0)
}

func (s *BlockStatementContext) Statement() IStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStatementContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IStatementContext)
}

func (s *BlockStatementContext) FunctionDeclaration() IFunctionDeclarationContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFunctionDeclarationContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFunctionDeclarationContext)
}

func (s *BlockStatementContext) ClassDeclaration() IClassDeclarationContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IClassDeclarationContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IClassDeclarationContext)
}

func (s *BlockStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BlockStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BlockStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PlayScriptListener); ok {
		listenerT.EnterBlockStatement(s)
	}
}

func (s *BlockStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PlayScriptListener); ok {
		listenerT.ExitBlockStatement(s)
	}
}

func (s *BlockStatementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PlayScriptVisitor:
		return t.VisitBlockStatement(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *PlayScriptParser) BlockStatement() (localctx IBlockStatementContext) {
	localctx = NewBlockStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 58, PlayScriptParserRULE_blockStatement)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(309)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 31, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(303)
			p.VariableDeclarators()
		}
		{
			p.SetState(304)
			p.Match(PlayScriptParserSEMI)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(306)
			p.Statement()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(307)
			p.FunctionDeclaration()
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(308)
			p.ClassDeclaration()
		}

	}

	return localctx
}

// IStatementContext is an interface to support dynamic dispatch.
type IStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetIdentifierLabel returns the identifierLabel token.
	GetIdentifierLabel() antlr.Token

	// SetIdentifierLabel sets the identifierLabel token.
	SetIdentifierLabel(antlr.Token)

	// GetBlockLabel returns the blockLabel rule contexts.
	GetBlockLabel() IBlockContext

	// GetStatementExpression returns the statementExpression rule contexts.
	GetStatementExpression() IExpressionContext

	// SetBlockLabel sets the blockLabel rule contexts.
	SetBlockLabel(IBlockContext)

	// SetStatementExpression sets the statementExpression rule contexts.
	SetStatementExpression(IExpressionContext)

	// IsStatementContext differentiates from other interfaces.
	IsStatementContext()
}

type StatementContext struct {
	*antlr.BaseParserRuleContext
	parser              antlr.Parser
	blockLabel          IBlockContext
	statementExpression IExpressionContext
	identifierLabel     antlr.Token
}

func NewEmptyStatementContext() *StatementContext {
	var p = new(StatementContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = PlayScriptParserRULE_statement
	return p
}

func (*StatementContext) IsStatementContext() {}

func NewStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StatementContext {
	var p = new(StatementContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = PlayScriptParserRULE_statement

	return p
}

func (s *StatementContext) GetParser() antlr.Parser { return s.parser }

func (s *StatementContext) GetIdentifierLabel() antlr.Token { return s.identifierLabel }

func (s *StatementContext) SetIdentifierLabel(v antlr.Token) { s.identifierLabel = v }

func (s *StatementContext) GetBlockLabel() IBlockContext { return s.blockLabel }

func (s *StatementContext) GetStatementExpression() IExpressionContext { return s.statementExpression }

func (s *StatementContext) SetBlockLabel(v IBlockContext) { s.blockLabel = v }

func (s *StatementContext) SetStatementExpression(v IExpressionContext) { s.statementExpression = v }

func (s *StatementContext) Block() IBlockContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBlockContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
}

func (s *StatementContext) IF() antlr.TerminalNode {
	return s.GetToken(PlayScriptParserIF, 0)
}

func (s *StatementContext) ParExpression() IParExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IParExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IParExpressionContext)
}

func (s *StatementContext) AllStatement() []IStatementContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IStatementContext)(nil)).Elem())
	var tst = make([]IStatementContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IStatementContext)
		}
	}

	return tst
}

func (s *StatementContext) Statement(i int) IStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStatementContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IStatementContext)
}

func (s *StatementContext) ELSE() antlr.TerminalNode {
	return s.GetToken(PlayScriptParserELSE, 0)
}

func (s *StatementContext) FOR() antlr.TerminalNode {
	return s.GetToken(PlayScriptParserFOR, 0)
}

func (s *StatementContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(PlayScriptParserLPAREN, 0)
}

func (s *StatementContext) ForControl() IForControlContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IForControlContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IForControlContext)
}

func (s *StatementContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(PlayScriptParserRPAREN, 0)
}

func (s *StatementContext) WHILE() antlr.TerminalNode {
	return s.GetToken(PlayScriptParserWHILE, 0)
}

func (s *StatementContext) DO() antlr.TerminalNode {
	return s.GetToken(PlayScriptParserDO, 0)
}

func (s *StatementContext) SEMI() antlr.TerminalNode {
	return s.GetToken(PlayScriptParserSEMI, 0)
}

func (s *StatementContext) SWITCH() antlr.TerminalNode {
	return s.GetToken(PlayScriptParserSWITCH, 0)
}

func (s *StatementContext) LBRACE() antlr.TerminalNode {
	return s.GetToken(PlayScriptParserLBRACE, 0)
}

func (s *StatementContext) RBRACE() antlr.TerminalNode {
	return s.GetToken(PlayScriptParserRBRACE, 0)
}

func (s *StatementContext) AllSwitchBlockStatementGroup() []ISwitchBlockStatementGroupContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ISwitchBlockStatementGroupContext)(nil)).Elem())
	var tst = make([]ISwitchBlockStatementGroupContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ISwitchBlockStatementGroupContext)
		}
	}

	return tst
}

func (s *StatementContext) SwitchBlockStatementGroup(i int) ISwitchBlockStatementGroupContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISwitchBlockStatementGroupContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ISwitchBlockStatementGroupContext)
}

func (s *StatementContext) AllSwitchLabel() []ISwitchLabelContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ISwitchLabelContext)(nil)).Elem())
	var tst = make([]ISwitchLabelContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ISwitchLabelContext)
		}
	}

	return tst
}

func (s *StatementContext) SwitchLabel(i int) ISwitchLabelContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISwitchLabelContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ISwitchLabelContext)
}

func (s *StatementContext) RETURN() antlr.TerminalNode {
	return s.GetToken(PlayScriptParserRETURN, 0)
}

func (s *StatementContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *StatementContext) BREAK() antlr.TerminalNode {
	return s.GetToken(PlayScriptParserBREAK, 0)
}

func (s *StatementContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(PlayScriptParserIDENTIFIER, 0)
}

func (s *StatementContext) CONTINUE() antlr.TerminalNode {
	return s.GetToken(PlayScriptParserCONTINUE, 0)
}

func (s *StatementContext) COLON() antlr.TerminalNode {
	return s.GetToken(PlayScriptParserCOLON, 0)
}

func (s *StatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PlayScriptListener); ok {
		listenerT.EnterStatement(s)
	}
}

func (s *StatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PlayScriptListener); ok {
		listenerT.ExitStatement(s)
	}
}

func (s *StatementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PlayScriptVisitor:
		return t.VisitStatement(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *PlayScriptParser) Statement() (localctx IStatementContext) {
	localctx = NewStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 60, PlayScriptParserRULE_statement)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.SetState(374)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 38, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(311)

			var _x = p.Block()

			localctx.(*StatementContext).blockLabel = _x
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(312)
			p.Match(PlayScriptParserIF)
		}
		{
			p.SetState(313)
			p.ParExpression()
		}
		{
			p.SetState(314)
			p.Statement()
		}
		p.SetState(317)
		p.GetErrorHandler().Sync(p)

		if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 32, p.GetParserRuleContext()) == 1 {
			{
				p.SetState(315)
				p.Match(PlayScriptParserELSE)
			}
			{
				p.SetState(316)
				p.Statement()
			}

		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(319)
			p.Match(PlayScriptParserFOR)
		}
		{
			p.SetState(320)
			p.Match(PlayScriptParserLPAREN)
		}
		{
			p.SetState(321)
			p.ForControl()
		}
		{
			p.SetState(322)
			p.Match(PlayScriptParserRPAREN)
		}
		{
			p.SetState(323)
			p.Statement()
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(325)
			p.Match(PlayScriptParserWHILE)
		}
		{
			p.SetState(326)
			p.ParExpression()
		}
		{
			p.SetState(327)
			p.Statement()
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(329)
			p.Match(PlayScriptParserDO)
		}
		{
			p.SetState(330)
			p.Statement()
		}
		{
			p.SetState(331)
			p.Match(PlayScriptParserWHILE)
		}
		{
			p.SetState(332)
			p.ParExpression()
		}
		{
			p.SetState(333)
			p.Match(PlayScriptParserSEMI)
		}

	case 6:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(335)
			p.Match(PlayScriptParserSWITCH)
		}
		{
			p.SetState(336)
			p.ParExpression()
		}
		{
			p.SetState(337)
			p.Match(PlayScriptParserLBRACE)
		}
		p.SetState(341)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 33, p.GetParserRuleContext())

		for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
			if _alt == 1 {
				{
					p.SetState(338)
					p.SwitchBlockStatementGroup()
				}

			}
			p.SetState(343)
			p.GetErrorHandler().Sync(p)
			_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 33, p.GetParserRuleContext())
		}
		p.SetState(347)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == PlayScriptParserCASE || _la == PlayScriptParserDEFAULT {
			{
				p.SetState(344)
				p.SwitchLabel()
			}

			p.SetState(349)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(350)
			p.Match(PlayScriptParserRBRACE)
		}

	case 7:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(352)
			p.Match(PlayScriptParserRETURN)
		}
		p.SetState(354)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if (((_la-40)&-(0x1f+1)) == 0 && ((1<<uint((_la-40)))&((1<<(PlayScriptParserSUPER-40))|(1<<(PlayScriptParserTHIS-40))|(1<<(PlayScriptParserDECIMAL_LITERAL-40))|(1<<(PlayScriptParserHEX_LITERAL-40))|(1<<(PlayScriptParserOCT_LITERAL-40))|(1<<(PlayScriptParserBINARY_LITERAL-40))|(1<<(PlayScriptParserFLOAT_LITERAL-40))|(1<<(PlayScriptParserHEX_FLOAT_LITERAL-40))|(1<<(PlayScriptParserBOOL_LITERAL-40))|(1<<(PlayScriptParserCHAR_LITERAL-40))|(1<<(PlayScriptParserSTRING_LITERAL-40))|(1<<(PlayScriptParserNULL_LITERAL-40))|(1<<(PlayScriptParserLPAREN-40)))) != 0) || (((_la-75)&-(0x1f+1)) == 0 && ((1<<uint((_la-75)))&((1<<(PlayScriptParserBANG-75))|(1<<(PlayScriptParserTILDE-75))|(1<<(PlayScriptParserINC-75))|(1<<(PlayScriptParserDEC-75))|(1<<(PlayScriptParserADD-75))|(1<<(PlayScriptParserSUB-75)))) != 0) || _la == PlayScriptParserIDENTIFIER {
			{
				p.SetState(353)
				p.expression(0)
			}

		}
		{
			p.SetState(356)
			p.Match(PlayScriptParserSEMI)
		}

	case 8:
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(357)
			p.Match(PlayScriptParserBREAK)
		}
		p.SetState(359)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == PlayScriptParserIDENTIFIER {
			{
				p.SetState(358)
				p.Match(PlayScriptParserIDENTIFIER)
			}

		}
		{
			p.SetState(361)
			p.Match(PlayScriptParserSEMI)
		}

	case 9:
		p.EnterOuterAlt(localctx, 9)
		{
			p.SetState(362)
			p.Match(PlayScriptParserCONTINUE)
		}
		p.SetState(364)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == PlayScriptParserIDENTIFIER {
			{
				p.SetState(363)
				p.Match(PlayScriptParserIDENTIFIER)
			}

		}
		{
			p.SetState(366)
			p.Match(PlayScriptParserSEMI)
		}

	case 10:
		p.EnterOuterAlt(localctx, 10)
		{
			p.SetState(367)
			p.Match(PlayScriptParserSEMI)
		}

	case 11:
		p.EnterOuterAlt(localctx, 11)
		{
			p.SetState(368)

			var _x = p.expression(0)

			localctx.(*StatementContext).statementExpression = _x
		}
		{
			p.SetState(369)
			p.Match(PlayScriptParserSEMI)
		}

	case 12:
		p.EnterOuterAlt(localctx, 12)
		{
			p.SetState(371)

			var _m = p.Match(PlayScriptParserIDENTIFIER)

			localctx.(*StatementContext).identifierLabel = _m
		}
		{
			p.SetState(372)
			p.Match(PlayScriptParserCOLON)
		}
		{
			p.SetState(373)
			p.Statement()
		}

	}

	return localctx
}

// ISwitchBlockStatementGroupContext is an interface to support dynamic dispatch.
type ISwitchBlockStatementGroupContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSwitchBlockStatementGroupContext differentiates from other interfaces.
	IsSwitchBlockStatementGroupContext()
}

type SwitchBlockStatementGroupContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySwitchBlockStatementGroupContext() *SwitchBlockStatementGroupContext {
	var p = new(SwitchBlockStatementGroupContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = PlayScriptParserRULE_switchBlockStatementGroup
	return p
}

func (*SwitchBlockStatementGroupContext) IsSwitchBlockStatementGroupContext() {}

func NewSwitchBlockStatementGroupContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SwitchBlockStatementGroupContext {
	var p = new(SwitchBlockStatementGroupContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = PlayScriptParserRULE_switchBlockStatementGroup

	return p
}

func (s *SwitchBlockStatementGroupContext) GetParser() antlr.Parser { return s.parser }

func (s *SwitchBlockStatementGroupContext) AllSwitchLabel() []ISwitchLabelContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ISwitchLabelContext)(nil)).Elem())
	var tst = make([]ISwitchLabelContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ISwitchLabelContext)
		}
	}

	return tst
}

func (s *SwitchBlockStatementGroupContext) SwitchLabel(i int) ISwitchLabelContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISwitchLabelContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ISwitchLabelContext)
}

func (s *SwitchBlockStatementGroupContext) AllBlockStatement() []IBlockStatementContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IBlockStatementContext)(nil)).Elem())
	var tst = make([]IBlockStatementContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IBlockStatementContext)
		}
	}

	return tst
}

func (s *SwitchBlockStatementGroupContext) BlockStatement(i int) IBlockStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBlockStatementContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IBlockStatementContext)
}

func (s *SwitchBlockStatementGroupContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SwitchBlockStatementGroupContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SwitchBlockStatementGroupContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PlayScriptListener); ok {
		listenerT.EnterSwitchBlockStatementGroup(s)
	}
}

func (s *SwitchBlockStatementGroupContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PlayScriptListener); ok {
		listenerT.ExitSwitchBlockStatementGroup(s)
	}
}

func (s *SwitchBlockStatementGroupContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PlayScriptVisitor:
		return t.VisitSwitchBlockStatementGroup(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *PlayScriptParser) SwitchBlockStatementGroup() (localctx ISwitchBlockStatementGroupContext) {
	localctx = NewSwitchBlockStatementGroupContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 62, PlayScriptParserRULE_switchBlockStatementGroup)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(377)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == PlayScriptParserCASE || _la == PlayScriptParserDEFAULT {
		{
			p.SetState(376)
			p.SwitchLabel()
		}

		p.SetState(379)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	p.SetState(382)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<PlayScriptParserBOOLEAN)|(1<<PlayScriptParserBREAK)|(1<<PlayScriptParserBYTE)|(1<<PlayScriptParserCHAR)|(1<<PlayScriptParserCLASS)|(1<<PlayScriptParserCONTINUE)|(1<<PlayScriptParserDO)|(1<<PlayScriptParserDOUBLE)|(1<<PlayScriptParserFLOAT)|(1<<PlayScriptParserFOR)|(1<<PlayScriptParserIF)|(1<<PlayScriptParserINT)|(1<<PlayScriptParserLONG))) != 0) || (((_la-36)&-(0x1f+1)) == 0 && ((1<<uint((_la-36)))&((1<<(PlayScriptParserRETURN-36))|(1<<(PlayScriptParserSHORT-36))|(1<<(PlayScriptParserSUPER-36))|(1<<(PlayScriptParserSWITCH-36))|(1<<(PlayScriptParserTHIS-36))|(1<<(PlayScriptParserVOID-36))|(1<<(PlayScriptParserWHILE-36))|(1<<(PlayScriptParserFUNCTION-36))|(1<<(PlayScriptParserSTRING-36))|(1<<(PlayScriptParserDECIMAL_LITERAL-36))|(1<<(PlayScriptParserHEX_LITERAL-36))|(1<<(PlayScriptParserOCT_LITERAL-36))|(1<<(PlayScriptParserBINARY_LITERAL-36))|(1<<(PlayScriptParserFLOAT_LITERAL-36))|(1<<(PlayScriptParserHEX_FLOAT_LITERAL-36))|(1<<(PlayScriptParserBOOL_LITERAL-36))|(1<<(PlayScriptParserCHAR_LITERAL-36))|(1<<(PlayScriptParserSTRING_LITERAL-36))|(1<<(PlayScriptParserNULL_LITERAL-36))|(1<<(PlayScriptParserLPAREN-36))|(1<<(PlayScriptParserLBRACE-36)))) != 0) || (((_la-69)&-(0x1f+1)) == 0 && ((1<<uint((_la-69)))&((1<<(PlayScriptParserSEMI-69))|(1<<(PlayScriptParserBANG-69))|(1<<(PlayScriptParserTILDE-69))|(1<<(PlayScriptParserINC-69))|(1<<(PlayScriptParserDEC-69))|(1<<(PlayScriptParserADD-69))|(1<<(PlayScriptParserSUB-69)))) != 0) || _la == PlayScriptParserIDENTIFIER {
		{
			p.SetState(381)
			p.BlockStatement()
		}

		p.SetState(384)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// ISwitchLabelContext is an interface to support dynamic dispatch.
type ISwitchLabelContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetEnumConstantName returns the enumConstantName token.
	GetEnumConstantName() antlr.Token

	// SetEnumConstantName sets the enumConstantName token.
	SetEnumConstantName(antlr.Token)

	// GetConstantExpression returns the constantExpression rule contexts.
	GetConstantExpression() IExpressionContext

	// SetConstantExpression sets the constantExpression rule contexts.
	SetConstantExpression(IExpressionContext)

	// IsSwitchLabelContext differentiates from other interfaces.
	IsSwitchLabelContext()
}

type SwitchLabelContext struct {
	*antlr.BaseParserRuleContext
	parser             antlr.Parser
	constantExpression IExpressionContext
	enumConstantName   antlr.Token
}

func NewEmptySwitchLabelContext() *SwitchLabelContext {
	var p = new(SwitchLabelContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = PlayScriptParserRULE_switchLabel
	return p
}

func (*SwitchLabelContext) IsSwitchLabelContext() {}

func NewSwitchLabelContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SwitchLabelContext {
	var p = new(SwitchLabelContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = PlayScriptParserRULE_switchLabel

	return p
}

func (s *SwitchLabelContext) GetParser() antlr.Parser { return s.parser }

func (s *SwitchLabelContext) GetEnumConstantName() antlr.Token { return s.enumConstantName }

func (s *SwitchLabelContext) SetEnumConstantName(v antlr.Token) { s.enumConstantName = v }

func (s *SwitchLabelContext) GetConstantExpression() IExpressionContext { return s.constantExpression }

func (s *SwitchLabelContext) SetConstantExpression(v IExpressionContext) { s.constantExpression = v }

func (s *SwitchLabelContext) CASE() antlr.TerminalNode {
	return s.GetToken(PlayScriptParserCASE, 0)
}

func (s *SwitchLabelContext) COLON() antlr.TerminalNode {
	return s.GetToken(PlayScriptParserCOLON, 0)
}

func (s *SwitchLabelContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *SwitchLabelContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(PlayScriptParserIDENTIFIER, 0)
}

func (s *SwitchLabelContext) DEFAULT() antlr.TerminalNode {
	return s.GetToken(PlayScriptParserDEFAULT, 0)
}

func (s *SwitchLabelContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SwitchLabelContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SwitchLabelContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PlayScriptListener); ok {
		listenerT.EnterSwitchLabel(s)
	}
}

func (s *SwitchLabelContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PlayScriptListener); ok {
		listenerT.ExitSwitchLabel(s)
	}
}

func (s *SwitchLabelContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PlayScriptVisitor:
		return t.VisitSwitchLabel(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *PlayScriptParser) SwitchLabel() (localctx ISwitchLabelContext) {
	localctx = NewSwitchLabelContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 64, PlayScriptParserRULE_switchLabel)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(394)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case PlayScriptParserCASE:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(386)
			p.Match(PlayScriptParserCASE)
		}
		p.SetState(389)
		p.GetErrorHandler().Sync(p)
		switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 41, p.GetParserRuleContext()) {
		case 1:
			{
				p.SetState(387)

				var _x = p.expression(0)

				localctx.(*SwitchLabelContext).constantExpression = _x
			}

		case 2:
			{
				p.SetState(388)

				var _m = p.Match(PlayScriptParserIDENTIFIER)

				localctx.(*SwitchLabelContext).enumConstantName = _m
			}

		}
		{
			p.SetState(391)
			p.Match(PlayScriptParserCOLON)
		}

	case PlayScriptParserDEFAULT:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(392)
			p.Match(PlayScriptParserDEFAULT)
		}
		{
			p.SetState(393)
			p.Match(PlayScriptParserCOLON)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IForControlContext is an interface to support dynamic dispatch.
type IForControlContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetForUpdate returns the forUpdate rule contexts.
	GetForUpdate() IExpressionListContext

	// SetForUpdate sets the forUpdate rule contexts.
	SetForUpdate(IExpressionListContext)

	// IsForControlContext differentiates from other interfaces.
	IsForControlContext()
}

type ForControlContext struct {
	*antlr.BaseParserRuleContext
	parser    antlr.Parser
	forUpdate IExpressionListContext
}

func NewEmptyForControlContext() *ForControlContext {
	var p = new(ForControlContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = PlayScriptParserRULE_forControl
	return p
}

func (*ForControlContext) IsForControlContext() {}

func NewForControlContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ForControlContext {
	var p = new(ForControlContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = PlayScriptParserRULE_forControl

	return p
}

func (s *ForControlContext) GetParser() antlr.Parser { return s.parser }

func (s *ForControlContext) GetForUpdate() IExpressionListContext { return s.forUpdate }

func (s *ForControlContext) SetForUpdate(v IExpressionListContext) { s.forUpdate = v }

func (s *ForControlContext) EnhancedForControl() IEnhancedForControlContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IEnhancedForControlContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IEnhancedForControlContext)
}

func (s *ForControlContext) AllSEMI() []antlr.TerminalNode {
	return s.GetTokens(PlayScriptParserSEMI)
}

func (s *ForControlContext) SEMI(i int) antlr.TerminalNode {
	return s.GetToken(PlayScriptParserSEMI, i)
}

func (s *ForControlContext) ForInit() IForInitContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IForInitContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IForInitContext)
}

func (s *ForControlContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *ForControlContext) ExpressionList() IExpressionListContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionListContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionListContext)
}

func (s *ForControlContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ForControlContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ForControlContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PlayScriptListener); ok {
		listenerT.EnterForControl(s)
	}
}

func (s *ForControlContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PlayScriptListener); ok {
		listenerT.ExitForControl(s)
	}
}

func (s *ForControlContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PlayScriptVisitor:
		return t.VisitForControl(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *PlayScriptParser) ForControl() (localctx IForControlContext) {
	localctx = NewForControlContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 66, PlayScriptParserRULE_forControl)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(408)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 46, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(396)
			p.EnhancedForControl()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		p.SetState(398)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<PlayScriptParserBOOLEAN)|(1<<PlayScriptParserBYTE)|(1<<PlayScriptParserCHAR)|(1<<PlayScriptParserDOUBLE)|(1<<PlayScriptParserFLOAT)|(1<<PlayScriptParserINT)|(1<<PlayScriptParserLONG))) != 0) || (((_la-37)&-(0x1f+1)) == 0 && ((1<<uint((_la-37)))&((1<<(PlayScriptParserSHORT-37))|(1<<(PlayScriptParserSUPER-37))|(1<<(PlayScriptParserTHIS-37))|(1<<(PlayScriptParserFUNCTION-37))|(1<<(PlayScriptParserSTRING-37))|(1<<(PlayScriptParserDECIMAL_LITERAL-37))|(1<<(PlayScriptParserHEX_LITERAL-37))|(1<<(PlayScriptParserOCT_LITERAL-37))|(1<<(PlayScriptParserBINARY_LITERAL-37))|(1<<(PlayScriptParserFLOAT_LITERAL-37))|(1<<(PlayScriptParserHEX_FLOAT_LITERAL-37))|(1<<(PlayScriptParserBOOL_LITERAL-37))|(1<<(PlayScriptParserCHAR_LITERAL-37))|(1<<(PlayScriptParserSTRING_LITERAL-37))|(1<<(PlayScriptParserNULL_LITERAL-37))|(1<<(PlayScriptParserLPAREN-37)))) != 0) || (((_la-75)&-(0x1f+1)) == 0 && ((1<<uint((_la-75)))&((1<<(PlayScriptParserBANG-75))|(1<<(PlayScriptParserTILDE-75))|(1<<(PlayScriptParserINC-75))|(1<<(PlayScriptParserDEC-75))|(1<<(PlayScriptParserADD-75))|(1<<(PlayScriptParserSUB-75)))) != 0) || _la == PlayScriptParserIDENTIFIER {
			{
				p.SetState(397)
				p.ForInit()
			}

		}
		{
			p.SetState(400)
			p.Match(PlayScriptParserSEMI)
		}
		p.SetState(402)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if (((_la-40)&-(0x1f+1)) == 0 && ((1<<uint((_la-40)))&((1<<(PlayScriptParserSUPER-40))|(1<<(PlayScriptParserTHIS-40))|(1<<(PlayScriptParserDECIMAL_LITERAL-40))|(1<<(PlayScriptParserHEX_LITERAL-40))|(1<<(PlayScriptParserOCT_LITERAL-40))|(1<<(PlayScriptParserBINARY_LITERAL-40))|(1<<(PlayScriptParserFLOAT_LITERAL-40))|(1<<(PlayScriptParserHEX_FLOAT_LITERAL-40))|(1<<(PlayScriptParserBOOL_LITERAL-40))|(1<<(PlayScriptParserCHAR_LITERAL-40))|(1<<(PlayScriptParserSTRING_LITERAL-40))|(1<<(PlayScriptParserNULL_LITERAL-40))|(1<<(PlayScriptParserLPAREN-40)))) != 0) || (((_la-75)&-(0x1f+1)) == 0 && ((1<<uint((_la-75)))&((1<<(PlayScriptParserBANG-75))|(1<<(PlayScriptParserTILDE-75))|(1<<(PlayScriptParserINC-75))|(1<<(PlayScriptParserDEC-75))|(1<<(PlayScriptParserADD-75))|(1<<(PlayScriptParserSUB-75)))) != 0) || _la == PlayScriptParserIDENTIFIER {
			{
				p.SetState(401)
				p.expression(0)
			}

		}
		{
			p.SetState(404)
			p.Match(PlayScriptParserSEMI)
		}
		p.SetState(406)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if (((_la-40)&-(0x1f+1)) == 0 && ((1<<uint((_la-40)))&((1<<(PlayScriptParserSUPER-40))|(1<<(PlayScriptParserTHIS-40))|(1<<(PlayScriptParserDECIMAL_LITERAL-40))|(1<<(PlayScriptParserHEX_LITERAL-40))|(1<<(PlayScriptParserOCT_LITERAL-40))|(1<<(PlayScriptParserBINARY_LITERAL-40))|(1<<(PlayScriptParserFLOAT_LITERAL-40))|(1<<(PlayScriptParserHEX_FLOAT_LITERAL-40))|(1<<(PlayScriptParserBOOL_LITERAL-40))|(1<<(PlayScriptParserCHAR_LITERAL-40))|(1<<(PlayScriptParserSTRING_LITERAL-40))|(1<<(PlayScriptParserNULL_LITERAL-40))|(1<<(PlayScriptParserLPAREN-40)))) != 0) || (((_la-75)&-(0x1f+1)) == 0 && ((1<<uint((_la-75)))&((1<<(PlayScriptParserBANG-75))|(1<<(PlayScriptParserTILDE-75))|(1<<(PlayScriptParserINC-75))|(1<<(PlayScriptParserDEC-75))|(1<<(PlayScriptParserADD-75))|(1<<(PlayScriptParserSUB-75)))) != 0) || _la == PlayScriptParserIDENTIFIER {
			{
				p.SetState(405)

				var _x = p.ExpressionList()

				localctx.(*ForControlContext).forUpdate = _x
			}

		}

	}

	return localctx
}

// IForInitContext is an interface to support dynamic dispatch.
type IForInitContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsForInitContext differentiates from other interfaces.
	IsForInitContext()
}

type ForInitContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyForInitContext() *ForInitContext {
	var p = new(ForInitContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = PlayScriptParserRULE_forInit
	return p
}

func (*ForInitContext) IsForInitContext() {}

func NewForInitContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ForInitContext {
	var p = new(ForInitContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = PlayScriptParserRULE_forInit

	return p
}

func (s *ForInitContext) GetParser() antlr.Parser { return s.parser }

func (s *ForInitContext) VariableDeclarators() IVariableDeclaratorsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IVariableDeclaratorsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IVariableDeclaratorsContext)
}

func (s *ForInitContext) ExpressionList() IExpressionListContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionListContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionListContext)
}

func (s *ForInitContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ForInitContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ForInitContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PlayScriptListener); ok {
		listenerT.EnterForInit(s)
	}
}

func (s *ForInitContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PlayScriptListener); ok {
		listenerT.ExitForInit(s)
	}
}

func (s *ForInitContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PlayScriptVisitor:
		return t.VisitForInit(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *PlayScriptParser) ForInit() (localctx IForInitContext) {
	localctx = NewForInitContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 68, PlayScriptParserRULE_forInit)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(412)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 47, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(410)
			p.VariableDeclarators()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(411)
			p.ExpressionList()
		}

	}

	return localctx
}

// IEnhancedForControlContext is an interface to support dynamic dispatch.
type IEnhancedForControlContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsEnhancedForControlContext differentiates from other interfaces.
	IsEnhancedForControlContext()
}

type EnhancedForControlContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyEnhancedForControlContext() *EnhancedForControlContext {
	var p = new(EnhancedForControlContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = PlayScriptParserRULE_enhancedForControl
	return p
}

func (*EnhancedForControlContext) IsEnhancedForControlContext() {}

func NewEnhancedForControlContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *EnhancedForControlContext {
	var p = new(EnhancedForControlContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = PlayScriptParserRULE_enhancedForControl

	return p
}

func (s *EnhancedForControlContext) GetParser() antlr.Parser { return s.parser }

func (s *EnhancedForControlContext) TypeType() ITypeTypeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITypeTypeContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITypeTypeContext)
}

func (s *EnhancedForControlContext) VariableDeclaratorId() IVariableDeclaratorIdContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IVariableDeclaratorIdContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IVariableDeclaratorIdContext)
}

func (s *EnhancedForControlContext) COLON() antlr.TerminalNode {
	return s.GetToken(PlayScriptParserCOLON, 0)
}

func (s *EnhancedForControlContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *EnhancedForControlContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EnhancedForControlContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *EnhancedForControlContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PlayScriptListener); ok {
		listenerT.EnterEnhancedForControl(s)
	}
}

func (s *EnhancedForControlContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PlayScriptListener); ok {
		listenerT.ExitEnhancedForControl(s)
	}
}

func (s *EnhancedForControlContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PlayScriptVisitor:
		return t.VisitEnhancedForControl(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *PlayScriptParser) EnhancedForControl() (localctx IEnhancedForControlContext) {
	localctx = NewEnhancedForControlContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 70, PlayScriptParserRULE_enhancedForControl)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(414)
		p.TypeType()
	}
	{
		p.SetState(415)
		p.VariableDeclaratorId()
	}
	{
		p.SetState(416)
		p.Match(PlayScriptParserCOLON)
	}
	{
		p.SetState(417)
		p.expression(0)
	}

	return localctx
}

// IParExpressionContext is an interface to support dynamic dispatch.
type IParExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsParExpressionContext differentiates from other interfaces.
	IsParExpressionContext()
}

type ParExpressionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyParExpressionContext() *ParExpressionContext {
	var p = new(ParExpressionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = PlayScriptParserRULE_parExpression
	return p
}

func (*ParExpressionContext) IsParExpressionContext() {}

func NewParExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ParExpressionContext {
	var p = new(ParExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = PlayScriptParserRULE_parExpression

	return p
}

func (s *ParExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ParExpressionContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(PlayScriptParserLPAREN, 0)
}

func (s *ParExpressionContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *ParExpressionContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(PlayScriptParserRPAREN, 0)
}

func (s *ParExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ParExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ParExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PlayScriptListener); ok {
		listenerT.EnterParExpression(s)
	}
}

func (s *ParExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PlayScriptListener); ok {
		listenerT.ExitParExpression(s)
	}
}

func (s *ParExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PlayScriptVisitor:
		return t.VisitParExpression(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *PlayScriptParser) ParExpression() (localctx IParExpressionContext) {
	localctx = NewParExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 72, PlayScriptParserRULE_parExpression)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(419)
		p.Match(PlayScriptParserLPAREN)
	}
	{
		p.SetState(420)
		p.expression(0)
	}
	{
		p.SetState(421)
		p.Match(PlayScriptParserRPAREN)
	}

	return localctx
}

// IExpressionListContext is an interface to support dynamic dispatch.
type IExpressionListContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsExpressionListContext differentiates from other interfaces.
	IsExpressionListContext()
}

type ExpressionListContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExpressionListContext() *ExpressionListContext {
	var p = new(ExpressionListContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = PlayScriptParserRULE_expressionList
	return p
}

func (*ExpressionListContext) IsExpressionListContext() {}

func NewExpressionListContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionListContext {
	var p = new(ExpressionListContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = PlayScriptParserRULE_expressionList

	return p
}

func (s *ExpressionListContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionListContext) AllExpression() []IExpressionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExpressionContext)(nil)).Elem())
	var tst = make([]IExpressionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExpressionContext)
		}
	}

	return tst
}

func (s *ExpressionListContext) Expression(i int) IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *ExpressionListContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(PlayScriptParserCOMMA)
}

func (s *ExpressionListContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(PlayScriptParserCOMMA, i)
}

func (s *ExpressionListContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionListContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionListContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PlayScriptListener); ok {
		listenerT.EnterExpressionList(s)
	}
}

func (s *ExpressionListContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PlayScriptListener); ok {
		listenerT.ExitExpressionList(s)
	}
}

func (s *ExpressionListContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PlayScriptVisitor:
		return t.VisitExpressionList(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *PlayScriptParser) ExpressionList() (localctx IExpressionListContext) {
	localctx = NewExpressionListContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 74, PlayScriptParserRULE_expressionList)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(423)
		p.expression(0)
	}
	p.SetState(428)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == PlayScriptParserCOMMA {
		{
			p.SetState(424)
			p.Match(PlayScriptParserCOMMA)
		}
		{
			p.SetState(425)
			p.expression(0)
		}

		p.SetState(430)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IFunctionCallContext is an interface to support dynamic dispatch.
type IFunctionCallContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFunctionCallContext differentiates from other interfaces.
	IsFunctionCallContext()
}

type FunctionCallContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFunctionCallContext() *FunctionCallContext {
	var p = new(FunctionCallContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = PlayScriptParserRULE_functionCall
	return p
}

func (*FunctionCallContext) IsFunctionCallContext() {}

func NewFunctionCallContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FunctionCallContext {
	var p = new(FunctionCallContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = PlayScriptParserRULE_functionCall

	return p
}

func (s *FunctionCallContext) GetParser() antlr.Parser { return s.parser }

func (s *FunctionCallContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(PlayScriptParserIDENTIFIER, 0)
}

func (s *FunctionCallContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(PlayScriptParserLPAREN, 0)
}

func (s *FunctionCallContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(PlayScriptParserRPAREN, 0)
}

func (s *FunctionCallContext) ExpressionList() IExpressionListContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionListContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionListContext)
}

func (s *FunctionCallContext) THIS() antlr.TerminalNode {
	return s.GetToken(PlayScriptParserTHIS, 0)
}

func (s *FunctionCallContext) SUPER() antlr.TerminalNode {
	return s.GetToken(PlayScriptParserSUPER, 0)
}

func (s *FunctionCallContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FunctionCallContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FunctionCallContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PlayScriptListener); ok {
		listenerT.EnterFunctionCall(s)
	}
}

func (s *FunctionCallContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PlayScriptListener); ok {
		listenerT.ExitFunctionCall(s)
	}
}

func (s *FunctionCallContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PlayScriptVisitor:
		return t.VisitFunctionCall(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *PlayScriptParser) FunctionCall() (localctx IFunctionCallContext) {
	localctx = NewFunctionCallContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 76, PlayScriptParserRULE_functionCall)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(449)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case PlayScriptParserIDENTIFIER:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(431)
			p.Match(PlayScriptParserIDENTIFIER)
		}
		{
			p.SetState(432)
			p.Match(PlayScriptParserLPAREN)
		}
		p.SetState(434)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if (((_la-40)&-(0x1f+1)) == 0 && ((1<<uint((_la-40)))&((1<<(PlayScriptParserSUPER-40))|(1<<(PlayScriptParserTHIS-40))|(1<<(PlayScriptParserDECIMAL_LITERAL-40))|(1<<(PlayScriptParserHEX_LITERAL-40))|(1<<(PlayScriptParserOCT_LITERAL-40))|(1<<(PlayScriptParserBINARY_LITERAL-40))|(1<<(PlayScriptParserFLOAT_LITERAL-40))|(1<<(PlayScriptParserHEX_FLOAT_LITERAL-40))|(1<<(PlayScriptParserBOOL_LITERAL-40))|(1<<(PlayScriptParserCHAR_LITERAL-40))|(1<<(PlayScriptParserSTRING_LITERAL-40))|(1<<(PlayScriptParserNULL_LITERAL-40))|(1<<(PlayScriptParserLPAREN-40)))) != 0) || (((_la-75)&-(0x1f+1)) == 0 && ((1<<uint((_la-75)))&((1<<(PlayScriptParserBANG-75))|(1<<(PlayScriptParserTILDE-75))|(1<<(PlayScriptParserINC-75))|(1<<(PlayScriptParserDEC-75))|(1<<(PlayScriptParserADD-75))|(1<<(PlayScriptParserSUB-75)))) != 0) || _la == PlayScriptParserIDENTIFIER {
			{
				p.SetState(433)
				p.ExpressionList()
			}

		}
		{
			p.SetState(436)
			p.Match(PlayScriptParserRPAREN)
		}

	case PlayScriptParserTHIS:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(437)
			p.Match(PlayScriptParserTHIS)
		}
		{
			p.SetState(438)
			p.Match(PlayScriptParserLPAREN)
		}
		p.SetState(440)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if (((_la-40)&-(0x1f+1)) == 0 && ((1<<uint((_la-40)))&((1<<(PlayScriptParserSUPER-40))|(1<<(PlayScriptParserTHIS-40))|(1<<(PlayScriptParserDECIMAL_LITERAL-40))|(1<<(PlayScriptParserHEX_LITERAL-40))|(1<<(PlayScriptParserOCT_LITERAL-40))|(1<<(PlayScriptParserBINARY_LITERAL-40))|(1<<(PlayScriptParserFLOAT_LITERAL-40))|(1<<(PlayScriptParserHEX_FLOAT_LITERAL-40))|(1<<(PlayScriptParserBOOL_LITERAL-40))|(1<<(PlayScriptParserCHAR_LITERAL-40))|(1<<(PlayScriptParserSTRING_LITERAL-40))|(1<<(PlayScriptParserNULL_LITERAL-40))|(1<<(PlayScriptParserLPAREN-40)))) != 0) || (((_la-75)&-(0x1f+1)) == 0 && ((1<<uint((_la-75)))&((1<<(PlayScriptParserBANG-75))|(1<<(PlayScriptParserTILDE-75))|(1<<(PlayScriptParserINC-75))|(1<<(PlayScriptParserDEC-75))|(1<<(PlayScriptParserADD-75))|(1<<(PlayScriptParserSUB-75)))) != 0) || _la == PlayScriptParserIDENTIFIER {
			{
				p.SetState(439)
				p.ExpressionList()
			}

		}
		{
			p.SetState(442)
			p.Match(PlayScriptParserRPAREN)
		}

	case PlayScriptParserSUPER:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(443)
			p.Match(PlayScriptParserSUPER)
		}
		{
			p.SetState(444)
			p.Match(PlayScriptParserLPAREN)
		}
		p.SetState(446)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if (((_la-40)&-(0x1f+1)) == 0 && ((1<<uint((_la-40)))&((1<<(PlayScriptParserSUPER-40))|(1<<(PlayScriptParserTHIS-40))|(1<<(PlayScriptParserDECIMAL_LITERAL-40))|(1<<(PlayScriptParserHEX_LITERAL-40))|(1<<(PlayScriptParserOCT_LITERAL-40))|(1<<(PlayScriptParserBINARY_LITERAL-40))|(1<<(PlayScriptParserFLOAT_LITERAL-40))|(1<<(PlayScriptParserHEX_FLOAT_LITERAL-40))|(1<<(PlayScriptParserBOOL_LITERAL-40))|(1<<(PlayScriptParserCHAR_LITERAL-40))|(1<<(PlayScriptParserSTRING_LITERAL-40))|(1<<(PlayScriptParserNULL_LITERAL-40))|(1<<(PlayScriptParserLPAREN-40)))) != 0) || (((_la-75)&-(0x1f+1)) == 0 && ((1<<uint((_la-75)))&((1<<(PlayScriptParserBANG-75))|(1<<(PlayScriptParserTILDE-75))|(1<<(PlayScriptParserINC-75))|(1<<(PlayScriptParserDEC-75))|(1<<(PlayScriptParserADD-75))|(1<<(PlayScriptParserSUB-75)))) != 0) || _la == PlayScriptParserIDENTIFIER {
			{
				p.SetState(445)
				p.ExpressionList()
			}

		}
		{
			p.SetState(448)
			p.Match(PlayScriptParserRPAREN)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IExpressionContext is an interface to support dynamic dispatch.
type IExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetPrefix returns the prefix token.
	GetPrefix() antlr.Token

	// GetBop returns the bop token.
	GetBop() antlr.Token

	// GetPostfix returns the postfix token.
	GetPostfix() antlr.Token

	// SetPrefix sets the prefix token.
	SetPrefix(antlr.Token)

	// SetBop sets the bop token.
	SetBop(antlr.Token)

	// SetPostfix sets the postfix token.
	SetPostfix(antlr.Token)

	// IsExpressionContext differentiates from other interfaces.
	IsExpressionContext()
}

type ExpressionContext struct {
	*antlr.BaseParserRuleContext
	parser  antlr.Parser
	prefix  antlr.Token
	bop     antlr.Token
	postfix antlr.Token
}

func NewEmptyExpressionContext() *ExpressionContext {
	var p = new(ExpressionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = PlayScriptParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = PlayScriptParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) GetPrefix() antlr.Token { return s.prefix }

func (s *ExpressionContext) GetBop() antlr.Token { return s.bop }

func (s *ExpressionContext) GetPostfix() antlr.Token { return s.postfix }

func (s *ExpressionContext) SetPrefix(v antlr.Token) { s.prefix = v }

func (s *ExpressionContext) SetBop(v antlr.Token) { s.bop = v }

func (s *ExpressionContext) SetPostfix(v antlr.Token) { s.postfix = v }

func (s *ExpressionContext) Primary() IPrimaryContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPrimaryContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IPrimaryContext)
}

func (s *ExpressionContext) FunctionCall() IFunctionCallContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFunctionCallContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFunctionCallContext)
}

func (s *ExpressionContext) AllExpression() []IExpressionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExpressionContext)(nil)).Elem())
	var tst = make([]IExpressionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExpressionContext)
		}
	}

	return tst
}

func (s *ExpressionContext) Expression(i int) IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *ExpressionContext) ADD() antlr.TerminalNode {
	return s.GetToken(PlayScriptParserADD, 0)
}

func (s *ExpressionContext) SUB() antlr.TerminalNode {
	return s.GetToken(PlayScriptParserSUB, 0)
}

func (s *ExpressionContext) INC() antlr.TerminalNode {
	return s.GetToken(PlayScriptParserINC, 0)
}

func (s *ExpressionContext) DEC() antlr.TerminalNode {
	return s.GetToken(PlayScriptParserDEC, 0)
}

func (s *ExpressionContext) TILDE() antlr.TerminalNode {
	return s.GetToken(PlayScriptParserTILDE, 0)
}

func (s *ExpressionContext) BANG() antlr.TerminalNode {
	return s.GetToken(PlayScriptParserBANG, 0)
}

func (s *ExpressionContext) MUL() antlr.TerminalNode {
	return s.GetToken(PlayScriptParserMUL, 0)
}

func (s *ExpressionContext) DIV() antlr.TerminalNode {
	return s.GetToken(PlayScriptParserDIV, 0)
}

func (s *ExpressionContext) MOD() antlr.TerminalNode {
	return s.GetToken(PlayScriptParserMOD, 0)
}

func (s *ExpressionContext) AllLT() []antlr.TerminalNode {
	return s.GetTokens(PlayScriptParserLT)
}

func (s *ExpressionContext) LT(i int) antlr.TerminalNode {
	return s.GetToken(PlayScriptParserLT, i)
}

func (s *ExpressionContext) AllGT() []antlr.TerminalNode {
	return s.GetTokens(PlayScriptParserGT)
}

func (s *ExpressionContext) GT(i int) antlr.TerminalNode {
	return s.GetToken(PlayScriptParserGT, i)
}

func (s *ExpressionContext) LE() antlr.TerminalNode {
	return s.GetToken(PlayScriptParserLE, 0)
}

func (s *ExpressionContext) GE() antlr.TerminalNode {
	return s.GetToken(PlayScriptParserGE, 0)
}

func (s *ExpressionContext) EQUAL() antlr.TerminalNode {
	return s.GetToken(PlayScriptParserEQUAL, 0)
}

func (s *ExpressionContext) NOTEQUAL() antlr.TerminalNode {
	return s.GetToken(PlayScriptParserNOTEQUAL, 0)
}

func (s *ExpressionContext) BITAND() antlr.TerminalNode {
	return s.GetToken(PlayScriptParserBITAND, 0)
}

func (s *ExpressionContext) CARET() antlr.TerminalNode {
	return s.GetToken(PlayScriptParserCARET, 0)
}

func (s *ExpressionContext) BITOR() antlr.TerminalNode {
	return s.GetToken(PlayScriptParserBITOR, 0)
}

func (s *ExpressionContext) AND() antlr.TerminalNode {
	return s.GetToken(PlayScriptParserAND, 0)
}

func (s *ExpressionContext) OR() antlr.TerminalNode {
	return s.GetToken(PlayScriptParserOR, 0)
}

func (s *ExpressionContext) COLON() antlr.TerminalNode {
	return s.GetToken(PlayScriptParserCOLON, 0)
}

func (s *ExpressionContext) QUESTION() antlr.TerminalNode {
	return s.GetToken(PlayScriptParserQUESTION, 0)
}

func (s *ExpressionContext) ASSIGN() antlr.TerminalNode {
	return s.GetToken(PlayScriptParserASSIGN, 0)
}

func (s *ExpressionContext) ADD_ASSIGN() antlr.TerminalNode {
	return s.GetToken(PlayScriptParserADD_ASSIGN, 0)
}

func (s *ExpressionContext) SUB_ASSIGN() antlr.TerminalNode {
	return s.GetToken(PlayScriptParserSUB_ASSIGN, 0)
}

func (s *ExpressionContext) MUL_ASSIGN() antlr.TerminalNode {
	return s.GetToken(PlayScriptParserMUL_ASSIGN, 0)
}

func (s *ExpressionContext) DIV_ASSIGN() antlr.TerminalNode {
	return s.GetToken(PlayScriptParserDIV_ASSIGN, 0)
}

func (s *ExpressionContext) AND_ASSIGN() antlr.TerminalNode {
	return s.GetToken(PlayScriptParserAND_ASSIGN, 0)
}

func (s *ExpressionContext) OR_ASSIGN() antlr.TerminalNode {
	return s.GetToken(PlayScriptParserOR_ASSIGN, 0)
}

func (s *ExpressionContext) XOR_ASSIGN() antlr.TerminalNode {
	return s.GetToken(PlayScriptParserXOR_ASSIGN, 0)
}

func (s *ExpressionContext) RSHIFT_ASSIGN() antlr.TerminalNode {
	return s.GetToken(PlayScriptParserRSHIFT_ASSIGN, 0)
}

func (s *ExpressionContext) URSHIFT_ASSIGN() antlr.TerminalNode {
	return s.GetToken(PlayScriptParserURSHIFT_ASSIGN, 0)
}

func (s *ExpressionContext) LSHIFT_ASSIGN() antlr.TerminalNode {
	return s.GetToken(PlayScriptParserLSHIFT_ASSIGN, 0)
}

func (s *ExpressionContext) MOD_ASSIGN() antlr.TerminalNode {
	return s.GetToken(PlayScriptParserMOD_ASSIGN, 0)
}

func (s *ExpressionContext) DOT() antlr.TerminalNode {
	return s.GetToken(PlayScriptParserDOT, 0)
}

func (s *ExpressionContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(PlayScriptParserIDENTIFIER, 0)
}

func (s *ExpressionContext) THIS() antlr.TerminalNode {
	return s.GetToken(PlayScriptParserTHIS, 0)
}

func (s *ExpressionContext) LBRACK() antlr.TerminalNode {
	return s.GetToken(PlayScriptParserLBRACK, 0)
}

func (s *ExpressionContext) RBRACK() antlr.TerminalNode {
	return s.GetToken(PlayScriptParserRBRACK, 0)
}

func (s *ExpressionContext) TypeType() ITypeTypeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITypeTypeContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITypeTypeContext)
}

func (s *ExpressionContext) INSTANCEOF() antlr.TerminalNode {
	return s.GetToken(PlayScriptParserINSTANCEOF, 0)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PlayScriptListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PlayScriptListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (s *ExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PlayScriptVisitor:
		return t.VisitExpression(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *PlayScriptParser) Expression() (localctx IExpressionContext) {
	return p.expression(0)
}

func (p *PlayScriptParser) expression(_p int) (localctx IExpressionContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()
	_parentState := p.GetState()
	localctx = NewExpressionContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IExpressionContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 78
	p.EnterRecursionRule(localctx, 78, PlayScriptParserRULE_expression, _p)
	var _la int

	defer func() {
		p.UnrollRecursionContexts(_parentctx)
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(458)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 53, p.GetParserRuleContext()) {
	case 1:
		{
			p.SetState(452)
			p.Primary()
		}

	case 2:
		{
			p.SetState(453)
			p.FunctionCall()
		}

	case 3:
		{
			p.SetState(454)

			var _lt = p.GetTokenStream().LT(1)

			localctx.(*ExpressionContext).prefix = _lt

			_la = p.GetTokenStream().LA(1)

			if !(((_la-85)&-(0x1f+1)) == 0 && ((1<<uint((_la-85)))&((1<<(PlayScriptParserINC-85))|(1<<(PlayScriptParserDEC-85))|(1<<(PlayScriptParserADD-85))|(1<<(PlayScriptParserSUB-85)))) != 0) {
				var _ri = p.GetErrorHandler().RecoverInline(p)

				localctx.(*ExpressionContext).prefix = _ri
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		{
			p.SetState(455)
			p.expression(15)
		}

	case 4:
		{
			p.SetState(456)

			var _lt = p.GetTokenStream().LT(1)

			localctx.(*ExpressionContext).prefix = _lt

			_la = p.GetTokenStream().LA(1)

			if !(_la == PlayScriptParserBANG || _la == PlayScriptParserTILDE) {
				var _ri = p.GetErrorHandler().RecoverInline(p)

				localctx.(*ExpressionContext).prefix = _ri
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		{
			p.SetState(457)
			p.expression(14)
		}

	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(526)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 57, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(524)
			p.GetErrorHandler().Sync(p)
			switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 56, p.GetParserRuleContext()) {
			case 1:
				localctx = NewExpressionContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, PlayScriptParserRULE_expression)
				p.SetState(460)

				if !(p.Precpred(p.GetParserRuleContext(), 13)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 13)", ""))
				}
				{
					p.SetState(461)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*ExpressionContext).bop = _lt

					_la = p.GetTokenStream().LA(1)

					if !(((_la-89)&-(0x1f+1)) == 0 && ((1<<uint((_la-89)))&((1<<(PlayScriptParserMUL-89))|(1<<(PlayScriptParserDIV-89))|(1<<(PlayScriptParserMOD-89)))) != 0) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*ExpressionContext).bop = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(462)
					p.expression(14)
				}

			case 2:
				localctx = NewExpressionContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, PlayScriptParserRULE_expression)
				p.SetState(463)

				if !(p.Precpred(p.GetParserRuleContext(), 12)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 12)", ""))
				}
				{
					p.SetState(464)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*ExpressionContext).bop = _lt

					_la = p.GetTokenStream().LA(1)

					if !(_la == PlayScriptParserADD || _la == PlayScriptParserSUB) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*ExpressionContext).bop = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(465)
					p.expression(13)
				}

			case 3:
				localctx = NewExpressionContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, PlayScriptParserRULE_expression)
				p.SetState(466)

				if !(p.Precpred(p.GetParserRuleContext(), 11)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 11)", ""))
				}
				p.SetState(474)
				p.GetErrorHandler().Sync(p)
				switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 54, p.GetParserRuleContext()) {
				case 1:
					{
						p.SetState(467)
						p.Match(PlayScriptParserLT)
					}
					{
						p.SetState(468)
						p.Match(PlayScriptParserLT)
					}

				case 2:
					{
						p.SetState(469)
						p.Match(PlayScriptParserGT)
					}
					{
						p.SetState(470)
						p.Match(PlayScriptParserGT)
					}
					{
						p.SetState(471)
						p.Match(PlayScriptParserGT)
					}

				case 3:
					{
						p.SetState(472)
						p.Match(PlayScriptParserGT)
					}
					{
						p.SetState(473)
						p.Match(PlayScriptParserGT)
					}

				}
				{
					p.SetState(476)
					p.expression(12)
				}

			case 4:
				localctx = NewExpressionContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, PlayScriptParserRULE_expression)
				p.SetState(477)

				if !(p.Precpred(p.GetParserRuleContext(), 10)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 10)", ""))
				}
				{
					p.SetState(478)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*ExpressionContext).bop = _lt

					_la = p.GetTokenStream().LA(1)

					if !(((_la-73)&-(0x1f+1)) == 0 && ((1<<uint((_la-73)))&((1<<(PlayScriptParserGT-73))|(1<<(PlayScriptParserLT-73))|(1<<(PlayScriptParserLE-73))|(1<<(PlayScriptParserGE-73)))) != 0) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*ExpressionContext).bop = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(479)
					p.expression(11)
				}

			case 5:
				localctx = NewExpressionContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, PlayScriptParserRULE_expression)
				p.SetState(480)

				if !(p.Precpred(p.GetParserRuleContext(), 8)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 8)", ""))
				}
				{
					p.SetState(481)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*ExpressionContext).bop = _lt

					_la = p.GetTokenStream().LA(1)

					if !(_la == PlayScriptParserEQUAL || _la == PlayScriptParserNOTEQUAL) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*ExpressionContext).bop = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(482)
					p.expression(9)
				}

			case 6:
				localctx = NewExpressionContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, PlayScriptParserRULE_expression)
				p.SetState(483)

				if !(p.Precpred(p.GetParserRuleContext(), 7)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 7)", ""))
				}
				{
					p.SetState(484)

					var _m = p.Match(PlayScriptParserBITAND)

					localctx.(*ExpressionContext).bop = _m
				}
				{
					p.SetState(485)
					p.expression(8)
				}

			case 7:
				localctx = NewExpressionContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, PlayScriptParserRULE_expression)
				p.SetState(486)

				if !(p.Precpred(p.GetParserRuleContext(), 6)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 6)", ""))
				}
				{
					p.SetState(487)

					var _m = p.Match(PlayScriptParserCARET)

					localctx.(*ExpressionContext).bop = _m
				}
				{
					p.SetState(488)
					p.expression(7)
				}

			case 8:
				localctx = NewExpressionContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, PlayScriptParserRULE_expression)
				p.SetState(489)

				if !(p.Precpred(p.GetParserRuleContext(), 5)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 5)", ""))
				}
				{
					p.SetState(490)

					var _m = p.Match(PlayScriptParserBITOR)

					localctx.(*ExpressionContext).bop = _m
				}
				{
					p.SetState(491)
					p.expression(6)
				}

			case 9:
				localctx = NewExpressionContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, PlayScriptParserRULE_expression)
				p.SetState(492)

				if !(p.Precpred(p.GetParserRuleContext(), 4)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 4)", ""))
				}
				{
					p.SetState(493)

					var _m = p.Match(PlayScriptParserAND)

					localctx.(*ExpressionContext).bop = _m
				}
				{
					p.SetState(494)
					p.expression(5)
				}

			case 10:
				localctx = NewExpressionContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, PlayScriptParserRULE_expression)
				p.SetState(495)

				if !(p.Precpred(p.GetParserRuleContext(), 3)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 3)", ""))
				}
				{
					p.SetState(496)

					var _m = p.Match(PlayScriptParserOR)

					localctx.(*ExpressionContext).bop = _m
				}
				{
					p.SetState(497)
					p.expression(4)
				}

			case 11:
				localctx = NewExpressionContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, PlayScriptParserRULE_expression)
				p.SetState(498)

				if !(p.Precpred(p.GetParserRuleContext(), 2)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 2)", ""))
				}
				{
					p.SetState(499)

					var _m = p.Match(PlayScriptParserQUESTION)

					localctx.(*ExpressionContext).bop = _m
				}
				{
					p.SetState(500)
					p.expression(0)
				}
				{
					p.SetState(501)
					p.Match(PlayScriptParserCOLON)
				}
				{
					p.SetState(502)
					p.expression(3)
				}

			case 12:
				localctx = NewExpressionContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, PlayScriptParserRULE_expression)
				p.SetState(504)

				if !(p.Precpred(p.GetParserRuleContext(), 1)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 1)", ""))
				}
				{
					p.SetState(505)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*ExpressionContext).bop = _lt

					_la = p.GetTokenStream().LA(1)

					if !((((_la-72)&-(0x1f+1)) == 0 && ((1<<uint((_la-72)))&((1<<(PlayScriptParserASSIGN-72))|(1<<(PlayScriptParserADD_ASSIGN-72))|(1<<(PlayScriptParserSUB_ASSIGN-72))|(1<<(PlayScriptParserMUL_ASSIGN-72))|(1<<(PlayScriptParserDIV_ASSIGN-72))|(1<<(PlayScriptParserAND_ASSIGN-72))|(1<<(PlayScriptParserOR_ASSIGN-72))|(1<<(PlayScriptParserXOR_ASSIGN-72))|(1<<(PlayScriptParserMOD_ASSIGN-72))|(1<<(PlayScriptParserLSHIFT_ASSIGN-72)))) != 0) || _la == PlayScriptParserRSHIFT_ASSIGN || _la == PlayScriptParserURSHIFT_ASSIGN) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*ExpressionContext).bop = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(506)
					p.expression(1)
				}

			case 13:
				localctx = NewExpressionContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, PlayScriptParserRULE_expression)
				p.SetState(507)

				if !(p.Precpred(p.GetParserRuleContext(), 19)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 19)", ""))
				}
				{
					p.SetState(508)

					var _m = p.Match(PlayScriptParserDOT)

					localctx.(*ExpressionContext).bop = _m
				}
				p.SetState(512)
				p.GetErrorHandler().Sync(p)
				switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 55, p.GetParserRuleContext()) {
				case 1:
					{
						p.SetState(509)
						p.Match(PlayScriptParserIDENTIFIER)
					}

				case 2:
					{
						p.SetState(510)
						p.FunctionCall()
					}

				case 3:
					{
						p.SetState(511)
						p.Match(PlayScriptParserTHIS)
					}

				}

			case 14:
				localctx = NewExpressionContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, PlayScriptParserRULE_expression)
				p.SetState(514)

				if !(p.Precpred(p.GetParserRuleContext(), 18)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 18)", ""))
				}
				{
					p.SetState(515)
					p.Match(PlayScriptParserLBRACK)
				}
				{
					p.SetState(516)
					p.expression(0)
				}
				{
					p.SetState(517)
					p.Match(PlayScriptParserRBRACK)
				}

			case 15:
				localctx = NewExpressionContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, PlayScriptParserRULE_expression)
				p.SetState(519)

				if !(p.Precpred(p.GetParserRuleContext(), 16)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 16)", ""))
				}
				{
					p.SetState(520)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*ExpressionContext).postfix = _lt

					_la = p.GetTokenStream().LA(1)

					if !(_la == PlayScriptParserINC || _la == PlayScriptParserDEC) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*ExpressionContext).postfix = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}

			case 16:
				localctx = NewExpressionContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, PlayScriptParserRULE_expression)
				p.SetState(521)

				if !(p.Precpred(p.GetParserRuleContext(), 9)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 9)", ""))
				}
				{
					p.SetState(522)

					var _m = p.Match(PlayScriptParserINSTANCEOF)

					localctx.(*ExpressionContext).bop = _m
				}
				{
					p.SetState(523)
					p.TypeType()
				}

			}

		}
		p.SetState(528)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 57, p.GetParserRuleContext())
	}

	return localctx
}

// IPrimaryContext is an interface to support dynamic dispatch.
type IPrimaryContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsPrimaryContext differentiates from other interfaces.
	IsPrimaryContext()
}

type PrimaryContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPrimaryContext() *PrimaryContext {
	var p = new(PrimaryContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = PlayScriptParserRULE_primary
	return p
}

func (*PrimaryContext) IsPrimaryContext() {}

func NewPrimaryContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PrimaryContext {
	var p = new(PrimaryContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = PlayScriptParserRULE_primary

	return p
}

func (s *PrimaryContext) GetParser() antlr.Parser { return s.parser }

func (s *PrimaryContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(PlayScriptParserLPAREN, 0)
}

func (s *PrimaryContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *PrimaryContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(PlayScriptParserRPAREN, 0)
}

func (s *PrimaryContext) THIS() antlr.TerminalNode {
	return s.GetToken(PlayScriptParserTHIS, 0)
}

func (s *PrimaryContext) SUPER() antlr.TerminalNode {
	return s.GetToken(PlayScriptParserSUPER, 0)
}

func (s *PrimaryContext) Literal() ILiteralContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ILiteralContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ILiteralContext)
}

func (s *PrimaryContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(PlayScriptParserIDENTIFIER, 0)
}

func (s *PrimaryContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PrimaryContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PrimaryContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PlayScriptListener); ok {
		listenerT.EnterPrimary(s)
	}
}

func (s *PrimaryContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PlayScriptListener); ok {
		listenerT.ExitPrimary(s)
	}
}

func (s *PrimaryContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PlayScriptVisitor:
		return t.VisitPrimary(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *PlayScriptParser) Primary() (localctx IPrimaryContext) {
	localctx = NewPrimaryContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 80, PlayScriptParserRULE_primary)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(537)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case PlayScriptParserLPAREN:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(529)
			p.Match(PlayScriptParserLPAREN)
		}
		{
			p.SetState(530)
			p.expression(0)
		}
		{
			p.SetState(531)
			p.Match(PlayScriptParserRPAREN)
		}

	case PlayScriptParserTHIS:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(533)
			p.Match(PlayScriptParserTHIS)
		}

	case PlayScriptParserSUPER:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(534)
			p.Match(PlayScriptParserSUPER)
		}

	case PlayScriptParserDECIMAL_LITERAL, PlayScriptParserHEX_LITERAL, PlayScriptParserOCT_LITERAL, PlayScriptParserBINARY_LITERAL, PlayScriptParserFLOAT_LITERAL, PlayScriptParserHEX_FLOAT_LITERAL, PlayScriptParserBOOL_LITERAL, PlayScriptParserCHAR_LITERAL, PlayScriptParserSTRING_LITERAL, PlayScriptParserNULL_LITERAL:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(535)
			p.Literal()
		}

	case PlayScriptParserIDENTIFIER:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(536)
			p.Match(PlayScriptParserIDENTIFIER)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// ITypeListContext is an interface to support dynamic dispatch.
type ITypeListContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTypeListContext differentiates from other interfaces.
	IsTypeListContext()
}

type TypeListContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTypeListContext() *TypeListContext {
	var p = new(TypeListContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = PlayScriptParserRULE_typeList
	return p
}

func (*TypeListContext) IsTypeListContext() {}

func NewTypeListContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TypeListContext {
	var p = new(TypeListContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = PlayScriptParserRULE_typeList

	return p
}

func (s *TypeListContext) GetParser() antlr.Parser { return s.parser }

func (s *TypeListContext) AllTypeType() []ITypeTypeContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ITypeTypeContext)(nil)).Elem())
	var tst = make([]ITypeTypeContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ITypeTypeContext)
		}
	}

	return tst
}

func (s *TypeListContext) TypeType(i int) ITypeTypeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITypeTypeContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ITypeTypeContext)
}

func (s *TypeListContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(PlayScriptParserCOMMA)
}

func (s *TypeListContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(PlayScriptParserCOMMA, i)
}

func (s *TypeListContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TypeListContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TypeListContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PlayScriptListener); ok {
		listenerT.EnterTypeList(s)
	}
}

func (s *TypeListContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PlayScriptListener); ok {
		listenerT.ExitTypeList(s)
	}
}

func (s *TypeListContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PlayScriptVisitor:
		return t.VisitTypeList(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *PlayScriptParser) TypeList() (localctx ITypeListContext) {
	localctx = NewTypeListContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 82, PlayScriptParserRULE_typeList)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(539)
		p.TypeType()
	}
	p.SetState(544)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == PlayScriptParserCOMMA {
		{
			p.SetState(540)
			p.Match(PlayScriptParserCOMMA)
		}
		{
			p.SetState(541)
			p.TypeType()
		}

		p.SetState(546)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// ITypeTypeContext is an interface to support dynamic dispatch.
type ITypeTypeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTypeTypeContext differentiates from other interfaces.
	IsTypeTypeContext()
}

type TypeTypeContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTypeTypeContext() *TypeTypeContext {
	var p = new(TypeTypeContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = PlayScriptParserRULE_typeType
	return p
}

func (*TypeTypeContext) IsTypeTypeContext() {}

func NewTypeTypeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TypeTypeContext {
	var p = new(TypeTypeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = PlayScriptParserRULE_typeType

	return p
}

func (s *TypeTypeContext) GetParser() antlr.Parser { return s.parser }

func (s *TypeTypeContext) ClassOrInterfaceType() IClassOrInterfaceTypeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IClassOrInterfaceTypeContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IClassOrInterfaceTypeContext)
}

func (s *TypeTypeContext) FunctionType() IFunctionTypeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFunctionTypeContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFunctionTypeContext)
}

func (s *TypeTypeContext) PrimitiveType() IPrimitiveTypeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPrimitiveTypeContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IPrimitiveTypeContext)
}

func (s *TypeTypeContext) AllLBRACK() []antlr.TerminalNode {
	return s.GetTokens(PlayScriptParserLBRACK)
}

func (s *TypeTypeContext) LBRACK(i int) antlr.TerminalNode {
	return s.GetToken(PlayScriptParserLBRACK, i)
}

func (s *TypeTypeContext) AllRBRACK() []antlr.TerminalNode {
	return s.GetTokens(PlayScriptParserRBRACK)
}

func (s *TypeTypeContext) RBRACK(i int) antlr.TerminalNode {
	return s.GetToken(PlayScriptParserRBRACK, i)
}

func (s *TypeTypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TypeTypeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TypeTypeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PlayScriptListener); ok {
		listenerT.EnterTypeType(s)
	}
}

func (s *TypeTypeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PlayScriptListener); ok {
		listenerT.ExitTypeType(s)
	}
}

func (s *TypeTypeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PlayScriptVisitor:
		return t.VisitTypeType(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *PlayScriptParser) TypeType() (localctx ITypeTypeContext) {
	localctx = NewTypeTypeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 84, PlayScriptParserRULE_typeType)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(550)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case PlayScriptParserIDENTIFIER:
		{
			p.SetState(547)
			p.ClassOrInterfaceType()
		}

	case PlayScriptParserFUNCTION:
		{
			p.SetState(548)
			p.FunctionType()
		}

	case PlayScriptParserBOOLEAN, PlayScriptParserBYTE, PlayScriptParserCHAR, PlayScriptParserDOUBLE, PlayScriptParserFLOAT, PlayScriptParserINT, PlayScriptParserLONG, PlayScriptParserSHORT, PlayScriptParserSTRING:
		{
			p.SetState(549)
			p.PrimitiveType()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}
	p.SetState(556)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 61, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(552)
				p.Match(PlayScriptParserLBRACK)
			}
			{
				p.SetState(553)
				p.Match(PlayScriptParserRBRACK)
			}

		}
		p.SetState(558)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 61, p.GetParserRuleContext())
	}

	return localctx
}

// IFunctionTypeContext is an interface to support dynamic dispatch.
type IFunctionTypeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFunctionTypeContext differentiates from other interfaces.
	IsFunctionTypeContext()
}

type FunctionTypeContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFunctionTypeContext() *FunctionTypeContext {
	var p = new(FunctionTypeContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = PlayScriptParserRULE_functionType
	return p
}

func (*FunctionTypeContext) IsFunctionTypeContext() {}

func NewFunctionTypeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FunctionTypeContext {
	var p = new(FunctionTypeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = PlayScriptParserRULE_functionType

	return p
}

func (s *FunctionTypeContext) GetParser() antlr.Parser { return s.parser }

func (s *FunctionTypeContext) FUNCTION() antlr.TerminalNode {
	return s.GetToken(PlayScriptParserFUNCTION, 0)
}

func (s *FunctionTypeContext) TypeTypeOrVoid() ITypeTypeOrVoidContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITypeTypeOrVoidContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITypeTypeOrVoidContext)
}

func (s *FunctionTypeContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(PlayScriptParserLPAREN, 0)
}

func (s *FunctionTypeContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(PlayScriptParserRPAREN, 0)
}

func (s *FunctionTypeContext) TypeList() ITypeListContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITypeListContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITypeListContext)
}

func (s *FunctionTypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FunctionTypeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FunctionTypeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PlayScriptListener); ok {
		listenerT.EnterFunctionType(s)
	}
}

func (s *FunctionTypeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PlayScriptListener); ok {
		listenerT.ExitFunctionType(s)
	}
}

func (s *FunctionTypeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PlayScriptVisitor:
		return t.VisitFunctionType(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *PlayScriptParser) FunctionType() (localctx IFunctionTypeContext) {
	localctx = NewFunctionTypeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 86, PlayScriptParserRULE_functionType)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(559)
		p.Match(PlayScriptParserFUNCTION)
	}
	{
		p.SetState(560)
		p.TypeTypeOrVoid()
	}
	{
		p.SetState(561)
		p.Match(PlayScriptParserLPAREN)
	}
	p.SetState(563)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<PlayScriptParserBOOLEAN)|(1<<PlayScriptParserBYTE)|(1<<PlayScriptParserCHAR)|(1<<PlayScriptParserDOUBLE)|(1<<PlayScriptParserFLOAT)|(1<<PlayScriptParserINT)|(1<<PlayScriptParserLONG))) != 0) || (((_la-37)&-(0x1f+1)) == 0 && ((1<<uint((_la-37)))&((1<<(PlayScriptParserSHORT-37))|(1<<(PlayScriptParserFUNCTION-37))|(1<<(PlayScriptParserSTRING-37)))) != 0) || _la == PlayScriptParserIDENTIFIER {
		{
			p.SetState(562)
			p.TypeList()
		}

	}
	{
		p.SetState(565)
		p.Match(PlayScriptParserRPAREN)
	}

	return localctx
}

// IPrimitiveTypeContext is an interface to support dynamic dispatch.
type IPrimitiveTypeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsPrimitiveTypeContext differentiates from other interfaces.
	IsPrimitiveTypeContext()
}

type PrimitiveTypeContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPrimitiveTypeContext() *PrimitiveTypeContext {
	var p = new(PrimitiveTypeContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = PlayScriptParserRULE_primitiveType
	return p
}

func (*PrimitiveTypeContext) IsPrimitiveTypeContext() {}

func NewPrimitiveTypeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PrimitiveTypeContext {
	var p = new(PrimitiveTypeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = PlayScriptParserRULE_primitiveType

	return p
}

func (s *PrimitiveTypeContext) GetParser() antlr.Parser { return s.parser }

func (s *PrimitiveTypeContext) BOOLEAN() antlr.TerminalNode {
	return s.GetToken(PlayScriptParserBOOLEAN, 0)
}

func (s *PrimitiveTypeContext) CHAR() antlr.TerminalNode {
	return s.GetToken(PlayScriptParserCHAR, 0)
}

func (s *PrimitiveTypeContext) BYTE() antlr.TerminalNode {
	return s.GetToken(PlayScriptParserBYTE, 0)
}

func (s *PrimitiveTypeContext) SHORT() antlr.TerminalNode {
	return s.GetToken(PlayScriptParserSHORT, 0)
}

func (s *PrimitiveTypeContext) INT() antlr.TerminalNode {
	return s.GetToken(PlayScriptParserINT, 0)
}

func (s *PrimitiveTypeContext) LONG() antlr.TerminalNode {
	return s.GetToken(PlayScriptParserLONG, 0)
}

func (s *PrimitiveTypeContext) FLOAT() antlr.TerminalNode {
	return s.GetToken(PlayScriptParserFLOAT, 0)
}

func (s *PrimitiveTypeContext) DOUBLE() antlr.TerminalNode {
	return s.GetToken(PlayScriptParserDOUBLE, 0)
}

func (s *PrimitiveTypeContext) STRING() antlr.TerminalNode {
	return s.GetToken(PlayScriptParserSTRING, 0)
}

func (s *PrimitiveTypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PrimitiveTypeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PrimitiveTypeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PlayScriptListener); ok {
		listenerT.EnterPrimitiveType(s)
	}
}

func (s *PrimitiveTypeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PlayScriptListener); ok {
		listenerT.ExitPrimitiveType(s)
	}
}

func (s *PrimitiveTypeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PlayScriptVisitor:
		return t.VisitPrimitiveType(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *PlayScriptParser) PrimitiveType() (localctx IPrimitiveTypeContext) {
	localctx = NewPrimitiveTypeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 88, PlayScriptParserRULE_primitiveType)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(567)
		_la = p.GetTokenStream().LA(1)

		if !((((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<PlayScriptParserBOOLEAN)|(1<<PlayScriptParserBYTE)|(1<<PlayScriptParserCHAR)|(1<<PlayScriptParserDOUBLE)|(1<<PlayScriptParserFLOAT)|(1<<PlayScriptParserINT)|(1<<PlayScriptParserLONG))) != 0) || _la == PlayScriptParserSHORT || _la == PlayScriptParserSTRING) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// ICreatorContext is an interface to support dynamic dispatch.
type ICreatorContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsCreatorContext differentiates from other interfaces.
	IsCreatorContext()
}

type CreatorContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCreatorContext() *CreatorContext {
	var p = new(CreatorContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = PlayScriptParserRULE_creator
	return p
}

func (*CreatorContext) IsCreatorContext() {}

func NewCreatorContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CreatorContext {
	var p = new(CreatorContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = PlayScriptParserRULE_creator

	return p
}

func (s *CreatorContext) GetParser() antlr.Parser { return s.parser }

func (s *CreatorContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(PlayScriptParserIDENTIFIER, 0)
}

func (s *CreatorContext) Arguments() IArgumentsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IArgumentsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IArgumentsContext)
}

func (s *CreatorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CreatorContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CreatorContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PlayScriptListener); ok {
		listenerT.EnterCreator(s)
	}
}

func (s *CreatorContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PlayScriptListener); ok {
		listenerT.ExitCreator(s)
	}
}

func (s *CreatorContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PlayScriptVisitor:
		return t.VisitCreator(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *PlayScriptParser) Creator() (localctx ICreatorContext) {
	localctx = NewCreatorContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 90, PlayScriptParserRULE_creator)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(569)
		p.Match(PlayScriptParserIDENTIFIER)
	}
	{
		p.SetState(570)
		p.Arguments()
	}

	return localctx
}

// ISuperSuffixContext is an interface to support dynamic dispatch.
type ISuperSuffixContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSuperSuffixContext differentiates from other interfaces.
	IsSuperSuffixContext()
}

type SuperSuffixContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySuperSuffixContext() *SuperSuffixContext {
	var p = new(SuperSuffixContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = PlayScriptParserRULE_superSuffix
	return p
}

func (*SuperSuffixContext) IsSuperSuffixContext() {}

func NewSuperSuffixContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SuperSuffixContext {
	var p = new(SuperSuffixContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = PlayScriptParserRULE_superSuffix

	return p
}

func (s *SuperSuffixContext) GetParser() antlr.Parser { return s.parser }

func (s *SuperSuffixContext) Arguments() IArgumentsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IArgumentsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IArgumentsContext)
}

func (s *SuperSuffixContext) DOT() antlr.TerminalNode {
	return s.GetToken(PlayScriptParserDOT, 0)
}

func (s *SuperSuffixContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(PlayScriptParserIDENTIFIER, 0)
}

func (s *SuperSuffixContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SuperSuffixContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SuperSuffixContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PlayScriptListener); ok {
		listenerT.EnterSuperSuffix(s)
	}
}

func (s *SuperSuffixContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PlayScriptListener); ok {
		listenerT.ExitSuperSuffix(s)
	}
}

func (s *SuperSuffixContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PlayScriptVisitor:
		return t.VisitSuperSuffix(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *PlayScriptParser) SuperSuffix() (localctx ISuperSuffixContext) {
	localctx = NewSuperSuffixContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 92, PlayScriptParserRULE_superSuffix)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(578)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case PlayScriptParserLPAREN:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(572)
			p.Arguments()
		}

	case PlayScriptParserDOT:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(573)
			p.Match(PlayScriptParserDOT)
		}
		{
			p.SetState(574)
			p.Match(PlayScriptParserIDENTIFIER)
		}
		p.SetState(576)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == PlayScriptParserLPAREN {
			{
				p.SetState(575)
				p.Arguments()
			}

		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IArgumentsContext is an interface to support dynamic dispatch.
type IArgumentsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsArgumentsContext differentiates from other interfaces.
	IsArgumentsContext()
}

type ArgumentsContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyArgumentsContext() *ArgumentsContext {
	var p = new(ArgumentsContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = PlayScriptParserRULE_arguments
	return p
}

func (*ArgumentsContext) IsArgumentsContext() {}

func NewArgumentsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ArgumentsContext {
	var p = new(ArgumentsContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = PlayScriptParserRULE_arguments

	return p
}

func (s *ArgumentsContext) GetParser() antlr.Parser { return s.parser }

func (s *ArgumentsContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(PlayScriptParserLPAREN, 0)
}

func (s *ArgumentsContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(PlayScriptParserRPAREN, 0)
}

func (s *ArgumentsContext) ExpressionList() IExpressionListContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionListContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionListContext)
}

func (s *ArgumentsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ArgumentsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ArgumentsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PlayScriptListener); ok {
		listenerT.EnterArguments(s)
	}
}

func (s *ArgumentsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PlayScriptListener); ok {
		listenerT.ExitArguments(s)
	}
}

func (s *ArgumentsContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PlayScriptVisitor:
		return t.VisitArguments(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *PlayScriptParser) Arguments() (localctx IArgumentsContext) {
	localctx = NewArgumentsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 94, PlayScriptParserRULE_arguments)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(580)
		p.Match(PlayScriptParserLPAREN)
	}
	p.SetState(582)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if (((_la-40)&-(0x1f+1)) == 0 && ((1<<uint((_la-40)))&((1<<(PlayScriptParserSUPER-40))|(1<<(PlayScriptParserTHIS-40))|(1<<(PlayScriptParserDECIMAL_LITERAL-40))|(1<<(PlayScriptParserHEX_LITERAL-40))|(1<<(PlayScriptParserOCT_LITERAL-40))|(1<<(PlayScriptParserBINARY_LITERAL-40))|(1<<(PlayScriptParserFLOAT_LITERAL-40))|(1<<(PlayScriptParserHEX_FLOAT_LITERAL-40))|(1<<(PlayScriptParserBOOL_LITERAL-40))|(1<<(PlayScriptParserCHAR_LITERAL-40))|(1<<(PlayScriptParserSTRING_LITERAL-40))|(1<<(PlayScriptParserNULL_LITERAL-40))|(1<<(PlayScriptParserLPAREN-40)))) != 0) || (((_la-75)&-(0x1f+1)) == 0 && ((1<<uint((_la-75)))&((1<<(PlayScriptParserBANG-75))|(1<<(PlayScriptParserTILDE-75))|(1<<(PlayScriptParserINC-75))|(1<<(PlayScriptParserDEC-75))|(1<<(PlayScriptParserADD-75))|(1<<(PlayScriptParserSUB-75)))) != 0) || _la == PlayScriptParserIDENTIFIER {
		{
			p.SetState(581)
			p.ExpressionList()
		}

	}
	{
		p.SetState(584)
		p.Match(PlayScriptParserRPAREN)
	}

	return localctx
}

func (p *PlayScriptParser) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
	switch ruleIndex {
	case 39:
		var t *ExpressionContext = nil
		if localctx != nil {
			t = localctx.(*ExpressionContext)
		}
		return p.Expression_Sempred(t, predIndex)

	default:
		panic("No predicate with index: " + fmt.Sprint(ruleIndex))
	}
}

func (p *PlayScriptParser) Expression_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 0:
		return p.Precpred(p.GetParserRuleContext(), 13)

	case 1:
		return p.Precpred(p.GetParserRuleContext(), 12)

	case 2:
		return p.Precpred(p.GetParserRuleContext(), 11)

	case 3:
		return p.Precpred(p.GetParserRuleContext(), 10)

	case 4:
		return p.Precpred(p.GetParserRuleContext(), 8)

	case 5:
		return p.Precpred(p.GetParserRuleContext(), 7)

	case 6:
		return p.Precpred(p.GetParserRuleContext(), 6)

	case 7:
		return p.Precpred(p.GetParserRuleContext(), 5)

	case 8:
		return p.Precpred(p.GetParserRuleContext(), 4)

	case 9:
		return p.Precpred(p.GetParserRuleContext(), 3)

	case 10:
		return p.Precpred(p.GetParserRuleContext(), 2)

	case 11:
		return p.Precpred(p.GetParserRuleContext(), 1)

	case 12:
		return p.Precpred(p.GetParserRuleContext(), 19)

	case 13:
		return p.Precpred(p.GetParserRuleContext(), 18)

	case 14:
		return p.Precpred(p.GetParserRuleContext(), 16)

	case 15:
		return p.Precpred(p.GetParserRuleContext(), 9)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}
