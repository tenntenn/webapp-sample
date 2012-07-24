define('main',
    [
        'task/main'
    ],
    (task)->
        taskListViewModel = new task.ListViewModel()
        taskListViewModel.load()
        ko.applyBindings(taskListViewModel)
)