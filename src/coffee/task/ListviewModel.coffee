define('task/ListViewModel',
    [
        'task/ViewModel'
    ],
    (ViewModel)->
        class ListViewModel
            constructor:()->
                @tasks = ko.observableArray()
                @selected = ko.observable(null)

            isSelected:(task)=>
                task is @selected

            select:(task)->
                if task
                    @selected(task)

            newTask:()=>
                task = new ViewModel()
                @tasks.push(task)
                task.save()
                @select(task)

            load:()=>
                @tasks.removeAll()
                $.ajax(
                    type:"GET"
                    url:"/task"
                    dataType: "JSON"
                    success:(data, dataType)=>
                        for model in data
                            task = new ViewModel(model)
                            @task.push(task)                            
                )

                if @tasks.length > 0
                    @select(@tasks()[0])
)