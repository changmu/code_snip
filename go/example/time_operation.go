defer func(t time.Time) { rlog.Debugf("time used:%v", time.Since(t)) }(time.Now())