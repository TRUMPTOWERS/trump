require('../css/app.css');
const angular = require('angular');
var app = angular.module('trump', []);

// Get the blings
app.factory('BlingService', ['$http', ($http) => {
    return {
        getList: () => {
            return $http.get('/getAll');
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
                <div class="no-blings" ng-if="!blings.length">
                    No Blings to be had
                </div>
                <div class="blings" ng-if="blings.length">
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
                (resp) => {
                    scope.blings = resp.data;
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
                <a ng-href="//{{blingz().name}}" class="bling-bling">
                    {{ blingz().name }}
                </a>
            </div>
        `
    };
}]);
