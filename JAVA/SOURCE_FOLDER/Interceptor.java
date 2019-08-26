package com.deltaone.common.interceptor;

import java.util.Enumeration;
import java.util.regex.Pattern;

import javax.servlet.http.HttpServletRequest;
import javax.servlet.http.HttpServletResponse;
import javax.servlet.http.HttpSession;

import org.apache.commons.lang3.StringUtils;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;
import org.springframework.beans.factory.annotation.Value;
import org.springframework.mobile.device.Device;
import org.springframework.mobile.device.DeviceResolver;
import org.springframework.mobile.device.LiteDeviceResolver;
import org.springframework.web.servlet.HandlerInterceptor;
import org.springframework.web.servlet.ModelAndView;

import com.deltaone.common.common.ReqData;

public class AdminInterceptor implements HandlerInterceptor {

    // 허용 IP대역폭
	@Value("${app.ip.bandwidth}")
	private String bandWidths;
	
	private Logger logger = LoggerFactory.getLogger("AdminInterceptorLogger");
	
	private final DeviceResolver deviceResolver;
	
	
	public AdminInterceptor() {
		this(new LiteDeviceResolver());
    }
    public AdminInterceptor(DeviceResolver deviceResolver) {
		this.deviceResolver = deviceResolver;}
	
        @Override
        public boolean preHandle(HttpServletRequest request, HttpServletResponse response, Object o) throws Exception {
            ReqData reqData = new ReqData(request);
		
            String ip = request.getRemoteAddr();
            Boolean isValid = false;
            String validIp = "^([1-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(\\.([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])){3}$";
            logger.debug("ADMIN 접근 감지  IP : [{}]", ip);
		
            logger.info("ip : {}", ip);
            if("0:0:0:0:0:0:0:1".equals(ip)){
                logger.debug("로컬 접근");
                isValid = true;
            }

            if("127.0.0.1".equals(ip)){
                logger.debug("로컬 접근");
                isValid = true;
            }
		
            if(!Pattern.matches(validIp, ip)){
                logger.debug("비허용 IP : 비정상  IP");isValid = false;
            }else{
                long longIp = ipToLong(ip);
                if(!isValid){
                    String[] bandwidth = StringUtils.deleteWhitespace(bandWidths).split(";");
                    String start_bandwidth = bandwidth[0];
                    String end_bandwidth = bandwidth[1];
                    if(longIp >= ipToLong(start_bandwidth) && longIp <= ipToLong(end_bandwidth)){
                        logger.debug("ADMIN 접근 허용된 IP(대역폭에 포함된) : [" + ip + "] 으로 접근");
                        isValid = true;
                    }
                }
            }
		
		
            if(!isValid){
                logger.debug("ADMIN 접근 거절");
                response.sendRedirect("/");
                return false;
            }
            
            // Time Check
            long startTime = System.currentTimeMillis();
            request.setAttribute("startTime", startTime);
    
            // Device Info (spring-mobile-device 1.1.3 use)
            Device device = this.deviceResolver.resolveDevice(request);
            request.setAttribute("device", (device.isMobile() ?  "MOBILE" : (device.isTablet() ? "TABLET" : "PC")));
		
            logger.info("----------------------------------- Request Parameter Info -------------------------------------------------");
            logger.info(reqData.getMap().toString());
		
            // Param Logging
		    logger.info("=================================== Request device Info    =================================================");
            logger.info("DEVICE: " + (device.isMobile() ?  "MOBILE" : (device.isTablet() ? "TABLET" : "PC")));
            logger.info("IP: " + request.getRemoteAddr());
            logger.info("----------------------------------- Request Header Info    -------------------------------------------------");
            Enumeration<?> headers = request.getHeaderNames();
            while (headers.hasMoreElements()) {
                String headerName = (String) headers.nextElement();
                logger.info(headerName + "=" + request.getHeader(headerName));
            }
            
            // Session Logging
		logger.info("=================================== Session  Info        ==================================================");
        HttpSession session = request.getSession();
        logger.info("SESSIONID = " + session.getId());
		Enumeration<?> sessions = session.getAttributeNames();
		while (sessions.hasMoreElements()) {
			String sessionName = (String) sessions.nextElement();
			logger.info(sessionName + "=" + session.getAttribute(sessionName));
		}
		logger.info("=================================== Session  Info End    ==================================================");
		
		
		Object adminInfo = session.getAttribute("adminInfo");
		if(adminInfo==null){
			logger.info("관리자 세션 정보 없음. 로그인 페이지로 이동.");
			response.sendRedirect("/admin/login.able");
			return false;
		}
		
		return true;
	}

    @Override
    public void postHandle(HttpServletRequest request, HttpServletResponse response, Object o, ModelAndView modelAndView) throws Exception {
		// Time Check
		long startTime = (Long)request.getAttribute("startTime");
		long endTime = System.currentTimeMillis();
		long executeTime = endTime - startTime;

		if (logger.isInfoEnabled()) {
			logger.info("-------------------------------------------------------------------------------------------");
			logger.info("요청 URL  : " + request.getRequestURL().toString()+" 응답시간: " + executeTime + "ms");
			logger.info("-------------------------------------------------------------------------------------------");
		}
	}

	@Override
	public void afterCompletion(HttpServletRequest request, HttpServletResponse response, Object o, Exception e) throws Exception {

	}
	
	private long ipToLong(String ip){
		String[] arrIp = ip.split("\\.");
		long longIp = (Long.parseLong(arrIp[0]) << 24) + (Long.parseLong(arrIp[1]) << 16) + (Long.parseLong(arrIp[2]) << 8) + Long.parseLong(arrIp[3]);
		return longIp;
	}
}


