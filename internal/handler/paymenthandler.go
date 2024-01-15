package handler

//func (h *Handler) checkout(w http.ResponseWriter, r *http.Request) {
//	switch r.Method {
//	case "POST":
//		input := &payment.Email{}
//
//		if err := json.NewDecoder(r.Body).Decode(input); err != nil {
//			ErrorHandler(w, http.StatusInternalServerError)
//			return
//		}
//
//		//service part
//		session, err := h.service.PaymenterService.Checkout(input.Email)
//		if err != nil {
//			log.Print(err)
//			log.Print("asd")
//			ErrorHandler(w, http.StatusInternalServerError)
//			return
//		}
//
//		w.Header().Set("Content-Type", "application/json")
//		if err = json.NewEncoder(w).Encode(session); err != nil {
//			log.Print(err)
//			ErrorHandler(w, http.StatusInternalServerError)
//			return
//		}
//	default:
//		ErrorHandler(w, http.StatusMethodNotAllowed)
//		return
//	}
//}
