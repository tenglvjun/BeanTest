function isRealNum(val){
    if(val === "" || val ==null){
        return false;
    }
    if(!isNaN(val)){
        return true;
    }else{
        return false;
    }
}     

function OnGetArithmetics(argument) {
	var range = $("input[id='question.range']").val();
	var count = $("input[id='question.count']").val();

	if (!isRealNum(range) || (range < 0)) {
		alert("请输入正确的计算区间格式");
		return false;
	}

	if (!isRealNum(count)) {
		alert("请输入正确的题数格式");
		return false;
	}

	$.ajax({
	  url: '/arithmetic/GetArithmetics',
	  type: 'POST',
	  dataType: 'json',
	  data: {
	  	range: range,
	  	count: count
	  },
	  success: function(data, textStatus, xhr) {
	  	if (data["code"] == 0) {
	  		alert(data["msg"]);
	  		return false
	  	}

	  	$("div[id='quetions.title']").empty();

		var title = "";

		title += "<div style=\"text-align: center;\">";
		title += "<h1>测试题目</h1>";
		title += "</div>";	

		$("div[id='quetions.title']").append(title);

	  	var questions = data["data"];

	  	var questions_left = new Array();
	  	var questions_right = new Array();

	  	for (var i = 0; i < questions.length; i++) {
	  		item = "<p>"+questions[i]+"</p>";
	  		if (i % 2 == 0) {
	  			questions_left.push(item);
	  		} else {
	  			questions_right.push(item);
	  		}
	  	}

	  	$("div[id='arithmetic.quetions_left']").empty();
	  	$("div[id='arithmetic.quetions_right']").empty();
	  	$("div[id='arithmetic.quetions_left']").append(questions_left);
	  	$("div[id='arithmetic.quetions_right']").append(questions_right);
	  	
	  	
	  	return true;
	  },
	  error: function(xhr, textStatus, errorThrown) {
	    alert("请求发送失败");
	    return false;
	  }
	});

	return false;
}

function OnPrint(argument) {
	// body...
	var print_id = argument;

	$(print_id).print({
    	globalStyles: true,
    	mediaPrint: false,
    	stylesheet: null,
    	noPrintSelector: ".no-print",
    	iframe: true,
    	append: null,
    	prepend: null,
    	manuallyCopyFormValues: true,
    	deferred: $.Deferred(),
    	timeout: 750,
    	title: null,
    	doctype: '<!doctype html>'
	});
}