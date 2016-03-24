require('../css/app.css');
const angular = require('angular');
var app = angular.module('trump', []);

// Get the blings
app.factory('BlingService', ['$q', ($q) => {
    return {
        getList: () => {
            // Stub out blings for now
            return $q.resolve([
                'http://donateToTrump.drumpf',
                'http://trumpforpresident.drumpf'
            ]);
        }
    };
}]);

// Show me the blings
app.directive('blings', ['BlingService', (BlingService) => {
    return {
        restrict: 'E',
        scope: {
            blings: '&'
        },
        template: `
            <div class="active-blings-wrapper" >
                <div class="active-blings__header">
                    ★★★  ACTIVE BLINGS  ★★★
                </div>
                <div class="blings">
                    <bling ng-repeat="b in blings track by $index"
                        index="$index + 1",
                        blingz="b">
                    </bling>
                </div>
            </div>
        `,
        link: (scope) => {
            scope.blings = [];
            BlingService.getList().then(
                (blings) => {
                    scope.blings = blings;
                }
            );
        }
    };
}]);

// Each Bling
app.directive('bling', [() => {
    return {
        restrict: 'E',
        scope: {
            index: '&',
            blingz: '&'
        },
        template: `
            <div class="blings__bling">
                <span class="bling">
                    {{ index() }}
                </span>
                <a ng-href="{{blingz()}}" class="bling-bling">
                    {{ blingz() }}
                </a>
            </div>
        `
    };
}]);
