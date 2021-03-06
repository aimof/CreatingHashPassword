const API = require('./api.js');
const APIResult = require('./api_result.js');
const Element = require('./view.js');
var api = new API();
var element = new Element();

var $ = require('jquery');

/*****************************/
//function for generate passphrase
//result: response json format
function generatePassphrase_resolved(result) {
	var dom = document.getElementById('result_In_Generated_passphrase')
	var jquery_dom = $("#result_In_Generated_passphrase")
	if( dom.value != result.result ) {
	    jquery_dom.hide()
	    dom.value = result.result
	    jquery_dom.fadeIn('normal');
	}
}
function generatePassphrase_reject(result) {
	alert(result)
}
//generate passphrase event
function generatePassphrase(){
	var title = document.getElementById('title_In_PassphraseSetting').value
	var keyphrase = document.getElementById('keyphrase_In_PassphraseSetting').value
	var algorithmSelect = document.getElementById('algorithmSelect_In_PassphraseSetting').value
	var extraInfo = document.getElementById('extraInfo_In_PassphraseSetting').value
	var length = Number(document.getElementById('maxLength_In_PassphraseSetting').value)
	var disableSymbol = ($("[name=UseSymbol_In_PassphraseSetting]").prop("checked") != true)
	api.generatePassphrase(title, keyphrase, algorithmSelect, extraInfo, length, disableSymbol, new APIResult(generatePassphrase_resolved, generatePassphrase_reject));
}

stoteActiveEvent("SubmitGeneratePassphrase", generatePassphrase);
/*****************************/
//function for save setting
//result: response json format
function savePassphraseInfo_resolved(result) {
    //Get settings after save it
    getPassphraseInfoAction()
	alert("Success to save passphrase setting")
}

function savePassphraseInfo_reject(result) {
	alert(result)
}
//save setting event
function savePassphraseSeedSetting(){
	var title = document.getElementById('title_In_PassphraseSetting').value
	var algorithmSelect = document.getElementById('algorithmSelect_In_PassphraseSetting').value
	var extraInfo = document.getElementById('extraInfo_In_PassphraseSetting').value
	var length = Number(document.getElementById('maxLength_In_PassphraseSetting').value)
	var disableSymbol = ($("[name=UseSymbol_In_PassphraseSetting]").prop("checked") != true)
	api.savePassphraseInfo(title, algorithmSelect, extraInfo, length, disableSymbol, new APIResult(savePassphraseInfo_resolved, savePassphraseInfo_reject));
}

stoteActiveEvent("SubmitSaveSetting", savePassphraseSeedSetting);
/*****************************/
//function for save setting
//result: response json format
function createUser_resolved(result) {
	alert("Success to save passphrase setting")
}
function createUser_reject(result) {
	alert(result)
}
//Create user event
function createUserEvent() {
	var user = document.getElementById('Username_In_UserAccount').value
	var pass = document.getElementById('LoginPassphrase_In_UserAccount').value
	api.createUser(user, pass, new APIResult(createUser_resolved, createUser_reject))
}
/*****************************/
//Update user event
function updateUser_resolved(result) {
	alert("Success to update passphrase setting")
}
function updateUser_reject(result) {
	alert(result)
}
//Update user event
function updateUserEvent() {
	var pass = document.getElementById('LoginPassphrase_In_UserAccount').value
	api.updateUser(pass, new APIResult(updateUser_resolved, updateUser_reject))
}

/*****************************/
//Delete user event
function deleteUser_resolved(result) {
	alert("Success to delete passphrase setting")
}
function deleteUser_reject(result) {
	alert(result)
}
//Delete user event
function deleteUserEvent() {
	api.deleteUser(new APIResult(deleteUser_resolved, deleteUser_reject))
}
function submitUserAccount(){
	if ( $('input[name=SelectOperation_In_UserAccount]:eq(0)').prop('checked') ) {
		createUserEvent();
	} else if ($('input[name=SelectOperation_In_UserAccount]:eq(1)').prop('checked')){
		updateUserEvent();
	} else {
		deleteUserEvent();
	}
}
stoteActiveEvent("SubmitUser_In_UserAccount", submitUserAccount);
/*****************************/
//Get setting event
function getPassphraseInfo_resolved(result) {
	$("#user_passphrase_settings_In_PassphraseSettings").empty()

	if(!Object.keys(result).length) {
		return
	}

	result.forEach(function(item, index, array) {
		title_id = item.title + '_In_PassphraseSetting'
		html = '<tr><td algorithm="'+item.algorithm+'" extra="' +item.seed+ '" maxlength="' + item.length + '" disable_symbol="' + item.disable_symbol + '" id="' + title_id + '"> '+ item.title + '</td>'
		$("#user_passphrase_settings_In_PassphraseSettings").append(html)

		//add event
		$(document).off('click', '[id="' + title_id + '"]')
		$(document).on('click', '[id="' + title_id + '"]', function(event) {
			title = event.currentTarget.innerText
			document.getElementById('title_In_PassphraseSetting').value = title
			title_element = document.getElementById(title + '_In_PassphraseSetting')
			document.getElementById('algorithmSelect_In_PassphraseSetting').value = title_element.getAttribute('algorithm')
			document.getElementById('extraInfo_In_PassphraseSetting').value = title_element.getAttribute('extra')
			document.getElementById('maxLength_In_PassphraseSetting').value = title_element.getAttribute('maxlength')
			document.getElementById('UseSymbol_In_PassphraseSetting').checked = (title_element.getAttribute('disable_symbol') != 'true')
		});
	})
}
function getPassphraseInfo_reject(result) {
	alert(result)
}
//Get setting event
function getPassphraseInfoAction() {
	api.getPassphraseInfo(new APIResult(getPassphraseInfo_resolved, getPassphraseInfo_reject));
}
stoteActiveEvent("SubmitGetSetting_In_PassphraseSettings", getPassphraseInfoAction);

/*****************************/
//Get setting event
function logout_resolved(result) {
}
function logout_reject(result) {
}
function logoutUser() {
	api.logout(new APIResult(logout_resolved, logout_reject));
	$("#user_passphrase_settings_In_PassphraseSettings").empty()
}
stoteActiveEvent("SubmitLogoutUser_In_UserAccount", logoutUser)

//Get setting event
function stoteActiveEvent(eventId, eventFunc) {
	$(document).on('click', '[id="' + eventId + '"]', function(){
		eventFunc()
		document.activeElement.blur()
	});
}

/*****************************/
//Set select event
function getMaxLength(algorithm) {
	switch(algorithm) {
		case "sha256":
		case "sha512_256":
		case "black2s_256":
		case "sha3_256":
			return [64,32,16];
		case "sha384":
		case "sha3_384":
		case "black2s_384":
			return [96,48,32,16];
		case "sha512":
		case "sha3_512":
		case "black2s_512":
			return [128,64,32,16];
	}
}
function updateMaxLengthField() {
	var algorithmSelect = document.getElementById('algorithmSelect_In_PassphraseSetting').value
	$("#maxLength_In_PassphraseSetting").empty()
	getMaxLength(algorithmSelect).forEach(function(length) {
		var html = "<option>" + length + "</option>"
		$("#maxLength_In_PassphraseSetting").append(html)
	})
}
$(document).on('change', '[id="algorithmSelect_In_PassphraseSetting"]', updateMaxLengthField)
window.onload = function() {
	updateMaxLengthField();
}
