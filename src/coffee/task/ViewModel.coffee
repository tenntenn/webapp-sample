define('task/ViewModel',
    [
        'utils/mapper'
    ],
    (mapper)->
        # A view model of Task
        class ViewModel
            constructor:(model)->
                
                @key = -1
                @name = ko.observable('New Task')
                @isDone = ko.observable(false)
                @deadline = ko.observable(new Date())
                @priority = ko.observable(0)

                mapper.map(@, model)

            save:=>
                    
                $.ajax(
                    type:"PUT"
                    url:"/task"
                    data:
                        task: $.toJSON(mapper.toModel(@))
                    dataType: "JSON"
                    success:(data, dataType)=>
                        @key = data.key
                )
)