import java.util.*;

/*
 * @lc app=leetcode id=138 lang=java
 *
 * [138] Copy List with Random Pointer
 */
class Node {
    int val;
    Node next;
    Node random;

    public Node(int val) {
        this.val = val;
        this.next = null;
        this.random = null;
    }
}

// @lc code=start
class Solution {
    // List<Node> oldList;
    List<Node> newList;
    Map<Node, Integer> oldMap;
    // Map<Node, Integer> newMap;

    public Node copyRandomList(Node head) {
        // oldList = new ArrayList<>();
        oldMap = new HashMap<>();
        // newMap = new HashMap<>();
        newList = new ArrayList<>();
        Node newHead = recurAddNode(head, 0);
        return newHead;
    }

    private Node recurAddNode(Node head, int index) {
        if (head == null) {
            // newList = new ArrayList<>(Collections.nCopies(index, null));
            return null;
        }
        // oldList.add(head);
        oldMap.put(head, index);
        Node newHead = new Node(head.val);
        newList.add(newHead);
        // newMap.put(newHead, index);
        newHead.next = recurAddNode(head.next, index + 1);
        if (head.random == null) {
            newHead.random = null;
        } else {
            int randomIndex = oldMap.get(head.random);
            newHead.random = newList.get(randomIndex);
        }
        return newHead;
    }

}
// @lc code=end

// Accepted
// 19/19 cases passed (0 ms)
// Your runtime beats 100 % of java submissions
// Your memory usage beats 96.07 % of java submissions (41.7 MB)