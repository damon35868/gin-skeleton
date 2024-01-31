package task

func Boot() {
	go func() {
		/**
		 * @description: 定时任务在此处注册
		 */
		NewOrderTask().Run()
	}()
}
