var __bind = function(fn, me){ return function(){ return fn.apply(me, arguments); }; };

define('task/ViewModel', ['utils/mapper'], function(mapper) {
  var ViewModel;
  return ViewModel = (function() {

    function ViewModel(model) {
      this.save = __bind(this.save, this);      this.key = -1;
      this.name = ko.observable('New Task');
      this.isDone = ko.observable(false);
      this.deadline = ko.observable(new Date());
      this.priority = ko.observable(0);
      mapper.map(this, model);
    }

    ViewModel.prototype.save = function() {
      var _this = this;
      return $.ajax({
        type: "PUT",
        url: "/task",
        data: {
          task: $.toJSON(mapper.toModel(this))
        },
        dataType: "JSON",
        success: function(data, dataType) {
          return _this.key = data.key;
        }
      });
    };

    return ViewModel;

  })();
});
