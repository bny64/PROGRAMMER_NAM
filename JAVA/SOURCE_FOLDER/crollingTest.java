package com.test;

import java.util.ArrayList;
import java.util.HashMap;
import java.util.List;
import java.util.Map;

import org.jsoup.Jsoup;
import org.jsoup.nodes.Document;
import org.jsoup.nodes.Element;
import org.jsoup.select.Elements;

public class CrollingTest {

	public static void A() {
		List<Map<String, String>> mConditions = new ArrayList<Map<String, String>>();

		try {
			String domain = "http://fok.funddoctor.co.kr";
			String url = "/board/aboard/board.jsp?name=fund_sihwang";
			Document doc = Jsoup.connect(domain + url).get();

			Elements elements = doc.getElementsByClass("fundCondition");
			elements = (Elements) elements.select("dl");
			for (Element element : elements) {

				Map<String, String> condition = new HashMap<String, String>();

				Elements a = element.select("a");
				Elements div = element.select("div.condition_excerpt");
				String imgPath = domain + element.select("img.fund_weather").attr("src");
				Elements span1 = element.select("span.condition_date");
				Elements span2 = element.select("span.condition_hit");
								
				String title = a.text();
				String content = div.text();
				String date = span1.text();
				String viewCount = span2.text();

				condition.put("title", title);
				condition.put("content", content);
				condition.put("imgPath", imgPath);
				condition.put("date", date);
				condition.put("viewCount", viewCount);

				if (title.contains("국내")) {
					condition.put("type", "국내");
				} else if (title.contains("해외")) {
					condition.put("type", "해외");
				} else {
					condition.put("type", "");
				}

				if (mConditions.size() < 1) {
					mConditions.add(condition);
				} else if (mConditions.size() > 1) {
					break;
				} else {
					for (int i = 0; i < mConditions.size(); i++) {
						if (condition.get("type").equals(mConditions.get(i).get("type"))) {
							continue;
						} else {
							mConditions.add(condition);
						}
					}
				}
			}

			System.out.println(mConditions);

		} catch (Exception e) {
			e.printStackTrace();
		}
	}

	public static void B(String type) {
		
		//1. 수익률 오름차, 2. 수익률 내림차, 3. 유입액 오름차, 4. 유입액 내림차 
		String list_orderItem = "22";		
		String fund_gb = "9";
		String search_type = "2";
		String list_orderBy = "";
		String ord_flag = "";
		
		switch(type) {
		case "1":
			list_orderBy = "S"; 
			ord_flag = "4";
			break;
		case "2":
			list_orderBy = "D"; 
			ord_flag = "4";
			break;
		case "3":
			list_orderBy = "S"; 
			ord_flag = "2";
			break;
		case "4":
			list_orderBy = "D"; 
			ord_flag = "2";
		}
		
		try {
			
			List<Map<String, String>> mStatistics = new ArrayList<Map<String, String>>();
			
			String domain = "http://fok.funddoctor.co.kr";
			String url = "/app/fund/zrntp_suikrt_all.jsp";
			
			Document doc = Jsoup.connect(domain+url)
					.data("list_orderItem",list_orderItem)
					.data("list_orderBy",list_orderBy)
					.data("ord_flag",ord_flag)
					.data("fund_gb",fund_gb)
					.data("search_type",search_type)
					.post();
			
			Elements elements = doc.select("tbody > tr");
			
			for(int i=0; i<elements.size(); i++) {
				
				if(mStatistics.size()>2) break;
				
				Map<String, String> statistics = new HashMap<String, String>();				
				
				Elements elements2 = elements.get(i).select("td");
				
				statistics.put("subject", elements2.get(1).text());
				statistics.put("foundCount", elements2.get(2).text());
				statistics.put("pureAsset", elements2.get(3).text());
				statistics.put("1monthIcm", elements2.get(4).text());
				statistics.put("3monthsIcm", elements2.get(5).text());
				statistics.put("6monthsIcm", elements2.get(6).text());
				statistics.put("1yearIcm", elements2.get(7).text());
				statistics.put("3yearsIcm", elements2.get(8).text());
				statistics.put("beginYearIcm", elements2.get(9).text());		
				
				mStatistics.add(statistics);	
				
			}
			
			for(Map a: mStatistics) {
				System.out.println(a);
			}
			
		}catch(Exception e) {
			e.printStackTrace();
		}
		
		
	}
	
	public static void main(String[] args) {
		// TODO Auto-generated method stub
		//A();
		System.out.println("\n");
		//B("1");
		System.out.println("\n");
		//B("2");
		System.out.println("\n");
		//B("3");
		System.out.println("\n");
		//B("4");
		System.out.println("\n");
		String a = "1212";
		System.out.println(a.split("10"));
		
	}

}
