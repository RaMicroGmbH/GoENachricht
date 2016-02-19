// Global variable 
var _b64data; 
var UsersList ;

$(document).ready(function () {  

	
	var Notification = window.Notification || window.mozNotification || window.webkitNotification;	
			
	/*---------------------------------------------------------------------
	|  Requesting Permission to activate Browser Notification            |		
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
				populateSubjectData(data);	
			},
			error: function(xhr, textStatus, error){
				console.log(xhr.statusText);
		        console.log(textStatus);
		       	console.log(error);
			}
		});		
		
	//TODO: disable drop-down if user starts typing in the Editor!
	// look for on-change event for html editor!
	
	//$('.dropdown-toggle').prop('disabled', true);
	/*
	$("#edt_ENachricht").focus( function(){
		$('#btn_subjectlist, #btn_caretvorlagen, #btn_auswahlliste').prop('disabled',true);
	});
	*/
	
	$('#edt_ENachricht').on('ightmleditorrendering', function (evt, ui) {                    
		$('#btn_subjectlist, #btn_caretvorlagen, #btn_auswahlliste').prop('disabled',true);
    });
	
	/*---------------------------------------------------------------------
	|    	Passing ENachricht Data to Server using AJAX POST  	     	   |		
	 ---------------------------------------------------------------------*/	
	
	$('#btn_main_save, #btn_main_send').on('click',function (e) { 

		var emsgContent = $("#edt_ENachricht").igHtmlEditor("getContent", "text");
		var _sofortnachrichtstatus = false;
		
		var currenttime = new Date();
				
		if($('#cb_sofortnachricht').is(':checked')){ 
			_sofortnachrichtstatus  = true;
		}
		
		if(_sofortnachrichtstatus){
			var instance = new Notification (
            	"Sofort Nachricht :" + "@" + currenttime , {
            	    body: emsgContent
            	}
        	);				
		}
		
		var _attachmentname = $('#in_attachmentname').val();				
		var enachrichtData  = { B64Data : _b64data,
								filename : currenttime,
								attachmentname : _attachmentname,
								enachricht  : emsgContent };
				
		// Avoid sending empty e-message to server ! 
		if ((emsgContent != "") && (emsgContent != "undefined")){
			
			$.ajax({
				type:"POST",
				url:"/updateENachrichtData",
				data: enachrichtData ,
				dataType: "json",
				success: function(data){
					console.log(data);  
					displayalert(data.Message, 'alert-success');	
					clearDataFields();						
				},
				error: function(xhr, textStatus, error){
					console.log(xhr.statusText);
			        console.log(textStatus);
			       	console.log(error);
				}
			});
		
		};		
	
	});		
   	
	/*---------------------------------------------------------------------
	|    Obtaining Base64data for attachement using Web API FileReader     |		
	 ---------------------------------------------------------------------*/
	
	$('#in_attachment').on('change',function(){
				
		var filesSelected = document.getElementById('in_attachment'	).files;
		
		if (filesSelected.length > 0){
			var fileToLoad = filesSelected[0];			
			var fileReader = new FileReader();
			var filename = fileToLoad.name;
			
			$('#in_attachmentname').val(filename);
			
			fileReader.onload = function(fileLoadedEvent) {
	          _b64data = fileLoadedEvent.target.result; // <--- data: base64				
			};
			fileReader.readAsDataURL(fileToLoad);		

		};
		
	});

	/*---------------------------------------------------------------------
	|   Requesting Auwahlliste Page from Server using AJAX POST            |		
	 ---------------------------------------------------------------------*/
	$('#btn_auswahlliste').on('click',function (e) {
		
		window.location.href = "/enachricht/auswahlliste";
		
		$.ajax({
			type:"GET",
			url:"/enachricht/auswahlliste",
			data: "" ,
			dataType: "json",
			success: function(data){
				//console.log(data);  				
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
	
	function populateSubjectData(data){
		var $li_openingTag = "<li id='select'><a class='listitem-subjects'>";
		var $li_closingTag = "</a></li>";
		$('#sublist').append($li_openingTag + "&nbsp;" + $li_closingTag);		
		$.each( data.IniData.Subject.Samples, function( key, value ) {                                    
            $('#sublist').append($li_openingTag + value + $li_closingTag);                       
        });		
	};	
	
	$('#sublist').on('click', 'li', function () { 	
		var val = $(this).text();		
		$('#btn_subjectlist').text(val);
	});
	
	function populateUserData(data){	
	
		$.each( data.IniData.Subject.Samples, function( key, value ) {		
		    if (value != ""){
				AddListItem("",value)
			}
		});
		
		var $li_openingTag = "<li id='select'><a class='listitem-users' data-value=";
		var $checkboxTag   = "><input class='cb_users' type='checkbox'/>&nbsp;";
		var $li_closingTag = "</a></li>";
				
		$('#userslist').append($li_openingTag + "Alle" + $checkboxTag + " Alle " + $li_closingTag);
		$.each( data.IniData.Users, function( key, value ) {
			$('#userslist').append($li_openingTag + value + $checkboxTag + value + $li_closingTag);									
      	});		
	};	
	
	UsersList = [];
	
	$('.listitem-users').on('click', 'a', function(event) {
		
		var $target = $( event.currentTarget ),
	       //val = $target.attr( 'data-value' ),
		   val = $(this).text(),
	       $inp = $target.find( 'input' ),
	       idx;	
		
		if ( ( idx = UsersList.indexOf( val ) ) > -1 ) {
	    	UsersList.splice( idx, 1 );
		  
			if($.trim(val) == "Alle"){
				$('.cb_users').prop('checked', false);
			}
			
	      setTimeout( function() { $inp.prop( 'checked', false ) }, 0);		
		   
	    } else {		 
	      	UsersList.push( val );
		  
			if($.trim(val) == "Alle") {
		    	$('.cb_users').prop('checked', true);
			}
			
	      	setTimeout( function() { $inp.prop( 'checked', true )}, 0);
	    }
		
	    $( event.target ).blur();	    
		
	    var selusers = UsersList.join(',').replace(/,/g, ';').split();	
			
	    if(UsersList.length === 0 ){	   		
			$('#btn_userslist').html('&nbsp;');
	    } else {
			$('#btn_userslist').text($(UsersList).get(-1)); 
	    }
		
		/*
		if($.trim(val) == "Alle"){
						
			$('.listitem-users ').find('li a').each( function(){
				console.log("Each User :" + $(this).text());
			});
		}
		*/
		
		//console.log("USERS_LIST :" + UsersList);
	    return false;
	});
		
	/*---------------------------------------------------------------------
	|   					Inserting a new List item 		               |		
	 ---------------------------------------------------------------------*/	
			
	$('#btn_add').on('click', function(){
		AddListItem("","")
	});		
	
	function AddListItem(ITX,LX){
		
		if (LX == ""){
					
			var XselId = $('#btn_subjectlist').text();
			console.log("selected option: " + XselId);
			if (XselId == " ") { 
				XselId="";			
			}
		}
		else{
			XselId = LX;
		}
	
		var totalitems = $('#subjectslist').children().length;
		var d = new Date();
		var date = d.getDay().toString() + d.getMonth().toString() + d.getFullYear().toString();
		var time = d.getHours().toString() + d.getMinutes().toString() + d.getSeconds().toString();
		var UNID = date + "" + time + "" + totalitems;
 		
		var $input   = '<input id="Ipt" class="in-sampleitem"/>';	
		var $btn_add = '<span id="PlusBtn" class="glyphicon glyphicon-plus text-success sp-glyph-add" title="Zeile einf&uuml;gen" font-size:1.5em;></span>';
		var $btn_rmv = '<span id="RemBtn" class="glyphicon glyphicon-remove sp-glyph-rmv" title="Zeile l&ouml;schen" font-size:1.5em;></span>';		
		var $li      = $(('<li id=' + UNID + ' class="list-group-item" style="height:46px;">') +  
					     ('<p id="tx" style="position:absolute; left:4px; top:11px;">' + totalitems + '.</p></li>')),
		
			$input   = $(($input) + ($btn_add) + ($btn_rmv),				  
				         { class: 'currentSubjectsList',     
						   type: 'text',
					       id: XselId}).val(XselId);

		$li.append($input);
          
		if (ITX == "") {$('#subjectslist').append($li);}
		if (ITX != "") {$('#subjectslist li').eq(ITX-1).before($li);} 

		var totalitems = $('#subjectslist').children().length;
		indexingListItem();
		//console.log("Total items: " + (totalitems-1) + "\n" + "Unique ID: " + UNID);		
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
		indexingListItem();
	});	
	
	/*---------------------------------------------------------------------
	|	    				Clear list 						    		   |		
	 ---------------------------------------------------------------------*/	
	
	$('#btn_auswahl_delete').on ('click',function(){
		$('#subjectslist li').remove();
	});

	/*---------------------------------------------------------------------
	|	    				Deleting a List item 			    			|		
	 ---------------------------------------------------------------------*/	
	$('#subjectslist').on('click', '#RemBtn',function(e){

		var Who = $(this).parent('li').attr('id');
		document.getElementById(Who).remove();
		console.log("Item ID = " + Who + " deleted");
	});	
	
	/*---------------------------------------------------------------------
	|    			Ziffernwahl 				      	                	|		
	 ---------------------------------------------------------------------*/	

	$('#cb_Ziffern').on('click',  function(){
		var Chx = $(this).is(':checked');
		indexingListItem();
	});

	/*---------------------------------------------------------------------
	|    	Enabling sorting 				      	            		   |		
	 ---------------------------------------------------------------------*/	
	
	$( '#subjectslist' ).sortable();
    //$( "#subjectslist" ).disableSelection()
	
	/*---------------------------------------------------------------------
	|    					Sorting List items     	           		   |		
	 ---------------------------------------------------------------------*/	
	
	$('#subjectslist').on('mousemove',function(){
		indexingListItem();
	});
	
	function indexingListItem() {
		var cb_status = $('#cb_Ziffern').is(':checked');
		var listitemsCount = $('#subjectslist').children().length;
		for ( i=0; i< (listitemsCount - 1) ; i++){
			if (cb_status == true) {$('#subjectslist li').eq(i).children('p').text((i+1)+".");}
			if (cb_status == false) {$('#subjectslist li').eq(i).children('p').text("");}
		}
	}

	/*---------------------------------------------------------------------
	|    				Save list to JSon				   			       |		
	 ---------------------------------------------------------------------*/

	$('#btn_auswahl_save').on('click',function() {
		
		auswahllist = [];
		indexingListItem();
		
		var Qty = $('#subjectslist').children().length;
		if (Qty <= 1) {return;}	
		
		for ( i=0; i< (Qty-1) ; i++){
			var IDX = $('#subjectslist li').eq(i).attr('id');
			var TXT = $('#subjectslist li').eq(i).children('input').val();
			var Pos = $('#subjectslist li').eq(i).children('p').text(); 
		
			console.log("Pos=" + Pos + " Line=" + (i+1) + " ID:" + IDX + "\n Text:" + TXT);
			
			auswahllist.push(TXT);
		}
		console.log("Auswahllist elements=" + auswahllist.length + "\n" + auswahllist);
	});
			
	/*---------------------------------------------------------------------
	|		Utility function for alert messages							   |		
	 ---------------------------------------------------------------------*/	
	function displayalert(message, alerttype){
		
		$('#alert_placeholder').append('<div id="alertdiv" class="alert ' + alerttype + '"><a class="close" data-dismiss="alert">×</a><span align="center">'+message+'</span></div>');
				
		window.setTimeout(function() {
		    $('#alertdiv').slideUp(1000, function(){
		        $(this).remove(); 
		    });
		}, 4000);
	}		
	
	function clearDataFields(){
		_b64data = " ";		
		$('#edt_ENachricht').igHtmlEditor("setContent", " ", "html");	
		$('#in_attachmentname').val("");
		UsersList = [];
	}
	
	/*---------------------------------------------------------------------
	|     Attachments : converted to base64 & send to server via AJAX      |		
	 ---------------------------------------------------------------------*/
	/*
	function encodeAttachmentAsURL(){

	    var filesSelected = document.getElementById('in_attachment'	).files;
		var _encB64 = "";
	    if (filesSelected.length > 0){
			
	        var fileToLoad = filesSelected[0];	
			var fname = fileToLoad.name;
			
			var fsize = parseFloat(fileToLoad.size) / parseFloat(1024 * 1024);
			
			console.log("FILE_SIZE:" + fsize + "MB");
			
			if(fsize > parseFloat(5 * 1024 *1024)){
				console.log("File exceeds the allowed Size!")
			}
			
			$('#in_attachmentname').val(fname);
			//document.getElementById('in_attachmentname').innerHTML = fname;
			//console.log("F_NAME :" + fname);
	        var fileReader = new FileReader();
						
	        fileReader.onload = function(fileLoadedEvent) {
	            var srcData = fileLoadedEvent.target.result; // <--- data: base64
	            var newAttachment = document.createElement('img');
	            newAttachment.src = srcData;
				
				_encB64 = newAttachment.outerHTML;
				
				//console.log("srcData :" + srcData);
				//console.log("ENC_B64 :" + _encB64);							
				//console.log("Ends processing attachment ...")				
				
				var encB64Data = { encB64 : srcData,
								   fn: fname}
				$.ajax({
					type:"POST",
					url:"/base64data",
					data: encB64Data ,
					dataType: "json",
					success: function(data){
						//--
					},
					error: function(xhr, textStatus, error){
						console.log(xhr.statusText);
				        console.log(textStatus);
				       	console.log(error);
					}
				});	
	        }
	        fileReader.readAsDataURL(fileToLoad);			
	    }
		
	};
	
	*/
});
