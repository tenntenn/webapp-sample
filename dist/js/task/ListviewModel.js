var __bind = function(fn, me){ return function(){ return fn.apply(me, arguments); }; };

define('task/ListViewModel', ['task/ViewModel'], function(ViewModel) {
  var ListViewModel;
  return ListViewModel = (function() {

    function ListViewModel() {
      this.load = __bind(this.load, this);
      this.newTask = __bind(this.newTask, this);
      this.isSelected = __bind(this.isSelected, this);      this.tasks = ko.observableArray();
      this.selected = ko.observable(null);
    }

    ListViewModel.prototype.isSelected = function(task) {
      return task === this.selected;
    };

    ListViewModel.prototype.select = function(task) {
      if (task) return this.selected(task);
    };

    ListViewModel.prototype.newTask = function() {
      var task;
      task = new ViewModel();
      this.tasks.push(task);
      task.save();
      return this.select(task);
    };

    ListViewModel.prototype.load = function() {
      var _this = this;
      this.tasks.removeAll();
      $.ajax({
        type: "GET",
        url: "/task",
        dataType: "JSON",
        success: function(data, dataType) {
          var model, task, _i, _len, _results;
          _results = [];
          for (_i = 0, _len = data.length; _i < _len; _i++) {
            model = data[_i];
            task = new ViewModel(model);
            _results.push(_this.task.push(task));
          }
          return _results;
        }
      });
      if (this.tasks.length > 0) return this.select(this.tasks()[0]);
    };

    return ListViewModel;

  })();
});
