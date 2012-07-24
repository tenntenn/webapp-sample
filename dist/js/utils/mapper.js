
define('utils/mapper', [], function() {
  return {
    map: function(viewModel, model) {
      var k, v, _results;
      if (model) {
        _results = [];
        for (k in model) {
          v = model[k];
          if (viewModel[k]) {
            _results.push(viewModel[k] = ko.observable(ko.utils.unwrapObservable(v)));
          } else {
            _results.push(void 0);
          }
        }
        return _results;
      }
    },
    toModel: function(viewModel) {
      var k, model, unwraped, v;
      if (viewModel) {
        model = {};
        for (k in viewModel) {
          v = viewModel[k];
          unwraped = ko.utils.unwrapObservable(v);
          if (viewModel.hasOwnProperty(k) && typeof unwraped !== "function") {
            model[k] = unwraped;
          }
        }
        return model;
      }
    }
  };
});
