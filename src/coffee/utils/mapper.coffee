define('utils/mapper',
    [
    ],
    ()->

        # map model's properties to viewModel
        map:(viewModel, model)->
            if model
                for k, v of model
                    viewModel[k] = ko.observable(ko.utils.unwrapObservable(v)) if viewModel[k]

        # create model object from viewModel
        toModel:(viewModel)->
            if viewModel
                model = {}
                for k, v of viewModel
                    unwraped = ko.utils.unwrapObservable(v)
                    if viewModel.hasOwnProperty(k) and typeof unwraped isnt "function"
                        model[k] = unwraped
                        
                return model
)