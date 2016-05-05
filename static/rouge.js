(function (window) {
  "use strict";
  var prefix = "/api/";
  
  function ParameterMissingError(paramName) {
    this.message = "Parameter Missing(" + paramName + ")";
  }
  function ValidationError() {
    this.message = "Validation Error";
  }
  function AjaxError(status, message) {
    this.message = "'" + status + "': " + message;
  }
  
  function  loadJson(url, validation) {
    return new Promise(function(resolve, reject) {
      fetch(url, {
        headers: {
          'Accept': 'application/json',
        }
      }).then(function(res){
        if(res.status != 200)
          reject(new AjaxError(res.status, res.statusText));
          
        return res.json();
      }).then(function (data) {
        if(typeof validation != 'function')
          resolve(data);
        else if(validation(data))
          resolve(data);
        else
          reject(new ValidationError())
      })
    });
  }
  function loadPaper(id) {
    return new Promise(function (resolve, reject) {
      if(!id){
        reject([new ParameterMissing("id")]);
      } else {
        var url = prefix + "papers/" + id;
        return loadJson(url, function(paper){
          return paper && paper.user && paper.content;
        });
      }
    })
  }
  
  window.Rouge = {
    loadPaper: loadPaper
  }
})(window)