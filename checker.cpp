#include <iostream>
#include <fstream>
#include <string>
#include <vector>
using namespace std;


void low(string &s) {
  for (char &c : s) {
    c = tolower(c);
  }
}
bool isvowl(char s) {
  if (s == 'a' || s == 'e' || s == 'i' || s == 'o' || s == 'u') { return true; }
  return false;
}

int takeit(char w1, char w2) {
  if (w1 == w2) { return 0; } 
  else if (isvowl(w1) && isvowl(w2)) { return 1; } 
  else if (!isvowl(w1) && !isvowl(w2)) { return 1; } 
  else { return 3; }
}

int opt (string word1, string word2) {
  int arr[word1.length()+1][word2.length()+1] = {0};
  for (int i = 0; i < word1.length()+1; i++) { arr[i][0] = 2 * i; }
  for (int i = 0; i < word2.length()+1; i++) { arr[0][i] = 2 * i; }

  for (int i = 1; i < word1.length()+1; i++) {
    for (int j = 1; j < word2.length()+1; j++) {
      int take = takeit(word1[i-1], word2[j-1]) + arr[i-1][j-1];
      int gap1 = 2 + arr[i-1][j];
      int gap2 = 2 + arr[i][j-1];
      int small = min(min(gap1, gap2), take);
      arr[i][j] = small;
    }
  }
  return arr[word1.length()][word2.length()];
}

struct block {
  string word;
  int val; 
};

vector<block> grabClosest(string word) {
  vector<block> word_list = {};
  block temp = {"", -1};
  for (int i = 0; i < 10; i++) { word_list.push_back(temp); };
  ifstream dict("dictionary.txt");
  string dict_word;

  while (getline(dict, dict_word)) {
    low(dict_word);
    int temp_val = opt(word, dict_word);
    
    for (int i = 0; i < 10; i++) {
      if (word_list[i].val > temp_val || word_list[i].val < 0) {
        for (int j = 9; j > i; j--) {
          word_list[j] = word_list[j-1];
        }
        word_list[i] = {dict_word, temp_val};
        break;
      }
    }
  }

  dict.close();
  return word_list;
}

int main(int argc, char** argv) {
  if (argc < 2) {
    cerr << "Error: no input provided" << endl;
    return 1;
  } else {

    string t = argv[1];
    low(t);

    vector<block> output = grabClosest(t);

    for (int i = 0; i < 10; i++) {
      cout << output[i].word << endl;
    }
    
    return 0;
  }
}
