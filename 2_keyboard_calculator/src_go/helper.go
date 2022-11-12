package main

import "time"

var interval interface{}
var session interface{}
var n interface{}
var keypresses map[interface{}]interface{}
var keypress_total_dict map[interface{}]interface{}
var keypresses_total []interface{}
var keypress_speed []interface{}

type Keys_Tracker struct {
	keypresses          map[interface{}]interface{}
	keypress_total_dict map[interface{}]interface{}
	start_time          float64
	interval            interface{}
	session             interface{}
}

func init() {
	interval = 5
	session = 60
	n = 3
	keypresses = map[interface{}]interface{}{}
	keypress_total_dict = map[interface{}]interface{}{}
	keypresses_total = []interface{}{}
	keypress_speed = []interface{}{}
}

func NewKeys_Tracker() (self *Keys_Tracker) {
	self = &Keys_Tracker{}
	self.keypresses = map[interface{}]interface{}{}
	self.keypress_total_dict = map[interface{}]interface{}{}
	self.start_time = float64(time.Now().UnixNano()) / 1000000000.0
	self.interval = interval
	self.session = session
	return
}

//store keypresses in dictionary
func (self *Keys_Tracker) store_keypresses(keypresses map[interface{}]interface{}) {
	for key, value := range keypresses {
		if _, ok := self.keypresses[key]; ok {
			self.keypresses[key] = self.keypresses[key].(int) + value.(int)
		} else {
			self.keypresses[key] = value
		}
	}
}

//store total keypresses in dictionary
func (self *Keys_Tracker) store_total_keypresses(keypresses map[interface{}]interface{}) {
	for key, value := range keypresses {
		if _, ok := self.keypress_total_dict[key]; ok {
			self.keypress_total_dict[key] = self.keypress_total_dict[key].(int) + value.(int)
		} else {
			self.keypress_total_dict[key] = value
		}
	}
}

//count how many keypresses in interval and store in dictionary
func (self *Keys_Tracker) count_keypresses(keypresses map[interface{}]interface{}) map[interface{}]interface{} {
	keypresses_count := map[interface{}]interface{}{}
	for key, value := range keypresses {
		if _, ok := keypresses_count[key]; ok {
			keypresses_count[key] = keypresses_count[key].(int) + value.(int)
		} else {
			keypresses_count[key] = value
		}
	}
	return keypresses_count
}

//return top n keypresses from store_keypresses dictionary
func (self *Keys_Tracker) get_top_keypresses(n interface{}) []interface{} {
	keypresses_total = []interface{}{}
	for key, value := range self.keypresses {
		keypresses_total = append(keypresses_total, []interface{}{key, value})
	}
	return keypresses_total
}

//return top n keypresses from store_total_keypresses dictionary
func (self *Keys_Tracker) get_top_total_keypresses(n interface{}) []interface{} {
	keypresses_total = []interface{}{}
	for key, value := range self.keypress_total_dict {
		keypresses_total = append(keypresses_total, []interface{}{key, value})
	}
	return keypresses_total
}

//calculate average keypresses speed per minute
func (self *Keys_Tracker) get_keypress_speed() []interface{} {
	keypress_speed = []interface{}{}
	for key, value := range self.keypresses {
		keypress_speed = append(keypress_speed, []interface{}{key, value.(int) / self.session.(int)})
	}
	return keypress_speed
}
