
define('main', ['task/main'], function(task) {
  var taskListViewModel;
  taskListViewModel = new task.ListViewModel();
  taskListViewModel.load();
  return ko.applyBindings(taskListViewModel);
});
