#include <iostream>
#include <vector>
#include <algorithm>

class Solution {
public:
    vector<int> kWeakestRows(vector<vector<int>>& mat, int k) {
        int m = mat.size();
        int n = mat[0].size();

        vector<pair<int, int>> strength;

        for (int ix = 0; ix < m; ix++) {
            int soldiers = count(mat[ix].begin(), mat[ix].end(), 1);

            strength.push_back({ix, soldiers});
        }

        sort(strength.begin(), strength.end(),
             [](const auto& a, const auto& b) {
                 return a.second == b.second ? a.first < b.first
                                             : a.second < b.second;
             });

        vector<int> ans;
        for (int ix = 0; ix < k; ix++) {
            ans.push_back(strength[ix].first);
        }  
        return ans;
    }
};