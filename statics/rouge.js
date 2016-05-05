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
  function qp(params) {
    var _params = Object.keys(params).filter(function(key){ return typeof(params[key]) == 'undefined'; });
    return _params.length > 0 
      ? "?" + _params.map(function(key) { return key + "=" + params[key];}).join("&")
      : "";
  }
  function  loadJson(url, validation) {
    return new Promise(function(resolve, reject) {
      fetch(url, {
        headers: {
          'Accept': 'application/json',
        }
      }).then(function(res){
        if(res.status != 200){
          reject(new AjaxError(res.status, res.statusText));
          return 
        } else
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
        resolve(url)
      }
    }).then(function(url){
      return loadJson(url, function(paper){
        return paper && paper.user && paper.content;
      });
    })
  }
  function loadPapers(params) {
    var url = prefix + "papers" + qp(params);
    return loadJson(url);
  }
  
  window.Rouge = {
    loadPaper: loadPaper,
    loadPapers: loadPapers
  }
})(window)