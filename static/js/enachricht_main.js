$(document).ready(function () {  

	
	var Notification = window.Notification || window.mozNotification || window.webkitNotification;	
			
	 /*---------------------------------------------------------------------
	 |	  Requesting Permission to activate Browser Notification            |		
	  ---------------------------------------------------------------------*/	
	 Notification.requestPermission(function (permission) {
     });
	
	/*---------------------------------------------------------------------
	|	   Default settings for IGNITE UI HTML Editor  					   |		
	 ---------------------------------------------------------------------*/
	$('#edt_ENachricht').igHtmlEditor({
		showCopyPasteToolbar: false,
		showInsertObjectToolbar: false,
		showTextToolbar: true,		
		width: "98%",
	});
	
	/*---------------------------------------------------------------------
	|    Requesting User default data from Server using AJAX      		   |		
	 ---------------------------------------------------------------------*/		
	
	$.ajax({
			type:"GET",
			url:"/getUserDefaultData",
			data: "" ,
			dataType: "json",
			success: function(data){
				populateUserData(data);					
			},
			error: function(xhr, textStatus, error){
				console.log(xhr.statusText);
		        console.log(textStatus);
		       	console.log(error);
			}
		});		
				
	/*---------------------------------------------------------------------
	|    	Passing ENachricht Data to Server using AJAX POST  	     	   |		
	 ---------------------------------------------------------------------*/
	
	
	$('#btn_main_save').on('click',function (e) { 
								
		var emsgContent = $("#edt_ENachricht").igHtmlEditor("getContent", "text");
		console.log("MSG :" + emsgContent);
		var ct = new Date();
		
		var instance = new Notification (
             "Sofort Nachricht :" + "@" + ct	, {
                body: emsgContent
            }
        );	
				
		var enachrichtData  = { enachricht  :  emsgContent };
		$.ajax({
			type:"POST",
			url:"/updateENachrichtData",
			data: enachrichtData ,
			dataType: "json",
			success: function(data){
				console.log(data);  				
			},
			error: function(xhr, textStatus, error){
				console.log(xhr.statusText);
		        console.log(textStatus);
		       	console.log(error);
			}
		});		
		
	});		
   
	/*---------------------------------------------------------------------
	|   Requesting Auwahlliste Page from Server using AJAX POST            |		
	 ---------------------------------------------------------------------*/
	$('#btn_auswahlliste').on('click',function (e) {
		console.log("Auswahliste fired!");
		window.location.href = "/enachricht/auswahlliste";
		
		var _vorlagen = "Test Vorlagen!";		
		var vorlagenData  = { vorlagen  :  _vorlagen };
		
		$.ajax({
			type:"POST",
			url:"/enachricht/auswahlliste",
			data: vorlagenData ,
			dataType: "json",
			success: function(data){
				console.log(data);  				
			},
			error: function(xhr, textStatus, error){
				console.log(xhr.statusText);
		        console.log(textStatus);
		       	console.log(error);
			}
		});	
	});
	
	
	$( "input[type=text]" ).focus(function() {
		  //$( this ).blur();
		alert("input is focused!!");
		});

	/*---------------------------------------------------------------------
	|    	Populating User Data under An-EPostkorb section		      	   |		
	 ---------------------------------------------------------------------*/
	
		//var uList = $('<li class="list-group-item">');
		//$(uList).html("Newly Added list..");
		//$('#subjectslist').append(uList);
	
	function populateUserData(data){
		console.log(data);  		

		$.each( data.IniData.Subject.Examples, function( key, value ) {		
		    if (value != ""){
				AddListItem("",value)
			}
		//		var $li = $('<li class="list-group-item"/>'),
		//		$input = $('<input style="border:none; width:95%; height:80%;" /> ', 
		//		{   class: 'currentSubjectsList',     
		//		type: 'text',
		//		id: value + '_' + key}).val(value);
		//		$li.append($input);
    	//      $("#subjectslist").append($li); 
		});

		/*
		$.each( data.IniData.Subject.Examples, function( key, value ) {		
		    if (value != ""){
				var uList = $('<li class="list-group-item">'+ value  + '</li>');			
				$('#subjectslist').append(uList);	
				$('#subjectslist').find('li:first').trigger('click');
			}	
		});	
		*/
		//----- TEST PURPOSE ONLY ----
		//----- <ul id="subjectslist" class="list-group test">      ("<ul id='ulDynamic' class='ulDynamic'></ul>");
		//$("#preSubjects").append("<ul id='list-group-item' class='list-group-item'></ul>");
       
		//$("#preSubjects").find("li:first").trigger("click");

		//------------------------------------------------------
	};

	/*---------------------------------------------------------------------
	|    				Inserting a new List item 			               |		
	 ---------------------------------------------------------------------*/	
	//$('#btn_add').on('click', function(){
		
	$('#btn_add').on('click', function(){
		AddListItem("","")
	});		

	function AddListItem(ITX,LX){
		if (LX == ""){
			var XselId=document.getElementById("Xselect").value;
			if (XselId == "--leer--") {XselId="";}
		}
		else 
		{
			XselId = LX;
		}
				
		var totalitems = $('#subjectslist').children().length;
		var d = new Date();
		var date = d.getDay().toString() + d.getMonth().toString() + d.getFullYear().toString();
		var time = d.getHours().toString() + d.getMinutes().toString() + d.getSeconds().toString();
		var UNID = date + "" + time + "" + totalitems;

		var $li = $(('<li id=' + UNID + ' class="list-group-item" style="height:46px;">') +  
					('<p id="tx" style="position:absolute; left:4px; top:11px;">' + totalitems + '.</p></li>')),
		$input = $(('<input id="Ipt" style="position:absolute; width:89%; height:50%; left:34px; top:10px; border:none;"/>') + 
				  ('<button id="RemBtn" style="position:absolute; width:25px; left:94%; top:8px;">-</button>') + 
				  ('<button id="PlusBtn" style="position:absolute; width:25px; left:97%; top:8px;">+</button>'), 
				{   class: 'currentSubjectsList',     
					type: 'text',
					id: XselId}).val(XselId);
			$li.append($input);
            
			if (ITX == "") {$('#subjectslist').append($li);}
			if (ITX != "") {$('#subjectslist li').eq(ITX-1).after($li);} 

		var totalitems = $('#subjectslist').children().length;
		sortingTable();
		console.log("Total items: " + (totalitems-1) + "\n" + "Unique ID: " + UNID);		
	};

	/*---------------------------------------------------------------------
	|    				Adding a List item with "+" button	   			   |		
	 ---------------------------------------------------------------------*/	
	$('#subjectslist').on('click', '#PlusBtn',function(e){
		
		var WX = $(this).parent('li').attr('id');
		var Qty = $('#subjectslist').children().length;
		for ( i=0; i< (Qty-1) ; i++){
			var IDX = $('#subjectslist li').eq(i).attr('id');
			if (IDX == WX) {
				console.log("IndexPosition= " + i);
				AddListItem(i+1,"");
				}
		console.log("Line=" + (i+1) + "\n ID:" + IDX + " - WHO:" + WX);
		}
		sortingTable()
	});	
	
	/*---------------------------------------------------------------------
	|	    				Clear list 						    		   |		
	 ---------------------------------------------------------------------*/	
	
	$('#btn_auswahl_delete').on ('click',function(){
		$('#subjectslist li').remove();
	});

	/*---------------------------------------------------------------------
	|	    				Deleting a List item 			    		   |		
	 ---------------------------------------------------------------------*/	
	$('#subjectslist').on('click', '#RemBtn',function(e){

		var Who = $(this).parent('li').attr('id');
		document.getElementById(Who).remove();
		console.log("Item ID = " + Who + " deleted");
	});	

	
	/*---------------------------------------------------------------------
	|    			Ziffernwahl 				      	                       |		
	 ---------------------------------------------------------------------*/	

	$('#cb_Ziffern').on('click',  function(){
		var Chx = $(this).is(':checked');
		sortingTable()
	});

	/*---------------------------------------------------------------------
	|    	Enabling sorting 				      	                       |		
	 ---------------------------------------------------------------------*/	
	
	$( "#subjectslist" ).sortable();
    //$( "#subjectslist" ).disableSelection()
	
	/*---------------------------------------------------------------------
	|    					Sorting List items     	                       |		
	 ---------------------------------------------------------------------*/	
	
	$('#subjectslist').on('mousemove',function(){
		sortingTable()
	});
	
	function sortingTable() {
		var Chx = $('#cb_Ziffern').is(':checked');
		var Qty = $('#subjectslist').children().length;
		for ( i=0; i< (Qty-1) ; i++){
			if (Chx == true) {$('#subjectslist li').eq(i).children('p').text((i+1)+".");}
			if (Chx == false) {$('#subjectslist li').eq(i).children('p').text("");}
			}
	}

	/*---------------------------------------------------------------------
	|    				Save list to JSon				             	   |		
	 ---------------------------------------------------------------------*/

	$('#btn_auswahl_save').on('click',function() {
		jsonObj = [];
		sortingTable()
		
		var Qty = $('#subjectslist').children().length;
		if (Qty <= 1) {return;}	
		
		for ( i=0; i< (Qty-1) ; i++){
			var IDX = $('#subjectslist li').eq(i).attr('id');
			var TXT = $('#subjectslist li').eq(i).children('input').val();
			var Pos = $('#subjectslist li').eq(i).children('p').text(); 
		
			console.log("Pos=" + Pos + " Line=" + (i+1) + " ID:" + IDX + "\n Text:" + TXT);
			jsonObj.push("ID=>" + IDX + "TX=>" + TXT);
		}
		console.log("Json elements=" + jsonObj.length + "\n" + jsonObj);
	});













	
});
