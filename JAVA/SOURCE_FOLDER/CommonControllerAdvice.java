package util;

import java.io.IOException;
import java.io.PrintWriter;
import java.io.StringWriter;
import java.util.ArrayList;
import java.util.List;
import java.util.regex.Matcher;
import java.util.regex.Pattern;

import javax.servlet.http.HttpServletRequest;
import javax.servlet.http.HttpServletResponse;

import org.springframework.http.HttpStatus;
import org.springframework.web.bind.annotation.ControllerAdvice;
import org.springframework.web.bind.annotation.ExceptionHandler;
import org.springframework.web.bind.annotation.ResponseStatus;
import org.springframework.web.servlet.ModelAndView;

@ControllerAdvice
public class CommonControllerAdvice {
	
	@ExceptionHandler(Exception.class)
	@ResponseStatus(HttpStatus.INTERNAL_SERVER_ERROR)
	public ModelAndView handleExcpetion(HttpServletRequest request, HttpServletResponse response, Exception e) throws IOException {
		
		String acceptHeader = request.getHeader("Accept");
		List<String> errList = generate(e);
		DocMake doc = new DocJson(request);
		
		if(acceptHeader.contains("application/json")) {
			if("E0001".equals(e.getLocalizedMessage())) {
				doc.setSetting("message");
			}else {
				doc.setSetting("message");
			}
			
			response.setContentType("application/json; charset=UTF-8");
			response.getWriter().append(doc.getDocument());
		}else {
			ModelAndView mnv = new ModelAndView();
			if("E0001".equals(e.getLocalizedMessage())){
				response.sendRedirect("/login.able");
			}else if("E0002".equals(e.getLocalizedMessage())){
				mnv.setViewName("error");
				mnv.addObject("errList","접근 불가 페이지");
			}else{
				mnv.setViewName("error");
				mnv.addObject("errList", errList);
			}

			String[] arrUrl = request.getRequestURL().toString().split("\\/");
			String viewName = arrUrl[arrUrl.length-1];
			String[] arrViewName = viewName.split("\\.");
			mnv.addObject("viewName", arrViewName[0]);

			return mnv;

		}
		
		return null;
		
		
	}
	
	private List<String> generate(Exception e) {
	    StringWriter writer = new StringWriter();
	    e.printStackTrace(new PrintWriter(writer));
	    String trace = writer.getBuffer().toString();
	    
	    Pattern tracePattern = Pattern.compile("\\s*at\\s+([\\w\\.$_]+)\\.([\\w$_]+)(\\\\(.*java)?:(\\\\d+)\\\\)(\\\\n|\\\\r\\\\n)");
	    Matcher matcher = tracePattern.matcher(trace);
	    List<String> rtnData = new ArrayList<String>();
	    while(matcher.find()) {
	        String className = matcher.group(1);
	        int lineNum = Integer.parseInt(matcher.group(4));
	        String[] arrName = className.split("\\.");
	        String data = String.valueOf(lineNum) + "-" + arrName[arrName.length-1];
	        rtnData.add(data);
	        
	        if(rtnData.size() > 1) {
	            break;
	        }
	    }
	    return rtnData;
	    
	}
	
}
