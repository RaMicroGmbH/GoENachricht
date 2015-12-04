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
	
	
	/*---------------------------------------------------------------------
	|    	Populating User Data under An-EPostkorb section		      	   |		
	 ---------------------------------------------------------------------*/
	function populateUserData(data){
		console.log(data);  		
		
		$.each( data.IniData.Users, function( key, value ) {					
			$('#userslist').append(new Option(value, value));				
		});	
		
		$.each( data.IniData.Subject.Examples, function( key, value ) {		
		    if (value != ""){
				var uList = $('<li  class="list-group-item">'+ value  + '</li>');			
				$('#subjectslist').append(uList);	
			}	
			
		});	
		
	};

	/*---------------------------------------------------------------------
	|    	Enabling sorting 				      	                       |		
	 ---------------------------------------------------------------------*/	
	
	$( "#subjectslist" ).sortable();
    $( "#subjectslist" ).disableSelection()
	

	/*---------------------------------------------------------------------
	|    	Inserting a new List item 			             	      	   |		
	 ---------------------------------------------------------------------*/	
	$('#btn_add').on('click', function(){
		
		var uList = $('<li  class="list-group-item">');
		$(uList).html("Newly Added list..");
		$('#subjectslist').append(uList);
		
	});

	
});
