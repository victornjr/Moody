// constructor(
//   // targetDiv,
//   formUrl, errorUrl) {
//     // this.targetDiv = document.getElementById("targetDiv");
//     // if (this.targetDiv == null) throw new Error(`Element ${targetDiv} was not found`);
//     this.formUrl = formUrl;
//     this.errorUrl = errorUrl;
// }

// getCookies = () => document.cookie
//     .split(";")
//     .map(c => c.trim())
//     .map(c => {
//         const tuple = c.split("=");
//         return { name: tuple[0], value: tuple[1] };
//     });

function hasFilledForm() {
  // const cookies = this.getCookies();
  // return cookies.includes(cookie => cookie.name === "filledForm" && cookie.value === "true");
  if(getCookie('f') == 'y'){
    return True;
  } else {
    return False;
  }
}

function getCookie(name) {
  var nameEQ = name + "=";
  var ca = document.cookie.split(';');
  for(var i=0;i < ca.length;i++) {
    var c = ca[i];
    while (c.charAt(0)==' ') c = c.substring(1,c.length);
    if (c.indexOf(nameEQ) == 0) return c.substring(nameEQ.length,c.length);
  }
  return null;
}

function setCookie(name,value,days) {
  var expires = "";
  if (days) {
    var date = new Date();
    date.setTime(date.getTime() + (days*24*60*60*1000));
    expires = "; expires=" + date.toUTCString();
  }
  document.cookie = name + "=" + (value || "")  + expires + "; path=/";
}

function eraseCookie(name) {
  document.cookie = name+'=; Max-Age=-99999999;';
}

module.exports = {
  hasFilledForm,
  getCookie,
  setCookie,
  eraseCookie
}
