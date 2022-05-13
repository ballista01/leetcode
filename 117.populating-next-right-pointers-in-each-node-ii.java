import java.util.LinkedList;
import java.util.Queue;
import org.w3c.dom.Node;
/*
 * @lc app=leetcode id=117 lang=java
 *
 * [117] Populating Next Right Pointers in Each Node II
 */
// @lc code=start
/*
// Definition for a Node.
class Node {
    public int val;
    public Node left;
    public Node right;
    public Node next;

    public Node() {}
    
    public Node(int _val) {
        val = _val;
    }

    public Node(int _val, Node _left, Node _right, Node _next) {
        val = _val;
        left = _left;
        right = _right;
        next = _next;
    }
};
*/

class Solution {
    public Node connect(Node root) {
        if (root == null)
            return null;
        LinkedList<Node> q = new LinkedList<Node>();
        q.add(root);
        Node tail = root;
        while (!q.isEmpty()) {
            Node node = q.remove();
            if (node.left != null) {
                q.add(node.left);
            }
            if (node.right != null) {
                q.add(node.right);
            }
            if (node == tail) {
                node.next = null;
                if (!q.isEmpty())
                    tail = q.getLast();
            } else {
                node.next = q.peek();
            }
        }
        return root;
    }
}
// @lc code=end

// Accepted
// 55/55 cases passed (1 ms)
// Your runtime beats 71.97 % of java submissions
// Your memory usage beats 96.34 % of java submissions (41.6 MB)
