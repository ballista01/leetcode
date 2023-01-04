import java.util.Arrays;

/*
 * @lc app=leetcode id=1105 lang=java
 *
 * [1105] Filling Bookcase Shelves
 */

// @lc code=start
class Solution {
    private int[][] books;
    private int shelfWidth;
    private int minHeight;
    // private List<Integer> shelves;

    public int minHeightShelves(int[][] books, int shelfWidth) {
        this.books = books;
        this.shelfWidth = shelfWidth;
        minHeight = Integer.MAX_VALUE;
        // for (int i = 0; i < books.length; i++)
        // System.out.println(Arrays.toString(books[i]));

        recurAddShelves(0, 0);

        return minHeight;
    }

    private void recurAddShelves(int start, int height) {

        System.out.printf("recurAddShelves: start=%d, height=%d\n", start, height);
        if (start >= books.length) {
            if (height < minHeight) {
                minHeight = height;
                return;
            }
        }
        int thisMaxHeight = 0;
        int maxEnd = start + getMaxBooksNum(start) - 1;
        System.out.printf(" maxEnd=%d\n", maxEnd);
        for (int i = start; i <= maxEnd; i++) {
            if (books[i][1] > thisMaxHeight) {
                thisMaxHeight = books[i][1];
            }
            recurAddShelves(i + 1, height + thisMaxHeight);
        }
    }

    private int getMaxBooksNum(int start) {
        int thisWidth = 0;
        int thisMaxHeight = 0;
        int i;
        for (i = start; i < books.length; i++) {
            thisWidth += books[i][0];
            if (thisWidth > shelfWidth) {
                return i - start;
            }
            if (books[i][1] > thisMaxHeight) {
                thisMaxHeight = books[i][1];
            }
        }
        return books.length - start;
    }
}
// @lc code=end
