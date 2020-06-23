package server

func (r *RechargeAPI) initRoutes() {
	r.router.Use(Log, Auth)
	r.router.HandleFunc("/show_plans", r.showplans)
	r.router.HandleFunc("/get_plans", r.getPlans)
	r.router.HandleFunc("/get_payment_link", r.getPaymentLink)
	r.router.HandleFunc("/check_payment", r.checkPayment)
}
