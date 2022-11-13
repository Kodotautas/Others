// package main

// import (
// 	"flag"
// )

// // add arguments
// func main() {
// 	interval := flag.Float64("interval", 1.0, "interval in seconds")
// 	session := flag.Float64("session", 60.0, "session in seconds")
// 	n := flag.Int("n", 5, "top n keypresses")
// 	flag.Parse()
// }

// var keypresses map[interface{}]interface{}
// var keypress_total_dict map[interface{}]interface{}
// var keypresses_total []interface{}
// var keypress_speed []interface{}

// // store keypresses in keypress dictionary
// func store_keypress(keypress interface{}) {
// 	if keypresses[keypress] == nil {
// 		keypresses[keypress] = 1
// 	} else {
// 		keypresses[keypress] = keypresses[keypress].(int) + 1
// 	}
// }

// // store total keypresses in dictionary
// func store_keypress_total(keypress interface{}) {
// 	if keypress_total_dict[keypress] == nil {
// 		keypress_total_dict[keypress] = 1
// 	} else {
// 		keypress_total_dict[keypress] = keypress_total_dict[keypress].(int) + 1
// 	}
// }

// // sum how many keypresses in interval and store in dictionary
// func count_keypresses() {
// 	for i := 0; i < len(keypresses_total); i++ {
// 		store_keypress(keypresses_total[i])
// 	}
// }

// // return top n keypresses from store_keypresses dictionary
// func top_n_keypresses(n int) []interface{} {
// 	var top_n []interface{}
// 	for i := 0; i < n; i++ {
// 		var max_key interface{}
// 		var max_value int
// 		for key, value := range keypresses {
// 			if value.(int) > max_value {
// 				max_value = value.(int)
// 				max_key = key
// 			}
// 		}
// 		top_n = append(top_n, max_key)
// 		delete(keypresses, max_key)
// 	}
// 	return top_n
// }

// // calculate average keypresses speed per minute
// func calculate_speed() float64 {
// 	var sum int
// 	for _, value := range keypress_speed {
// 		sum += value.(int)
// 	}
// 	return float64(sum) / int(*session) * 60
// }
