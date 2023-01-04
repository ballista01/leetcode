import java.util.*;

/*
 * @lc app=leetcode id=1090 lang=java
 *
 * [1090] Largest Values From Labels
 */

// @lc code=start
class Solution {
    public int largestValsFromLabels(int[] values, int[] labels, int numWanted, int useLimit) {
        List<Node> nodeList = new ArrayList<>();
        Map<Integer, Integer> labelLimit = new HashMap<>();
        for (int i = 0; i < values.length; i++) {
            nodeList.add(new Node(values[i], labels[i]));
            labelLimit.put(labels[i], useLimit);
        }
        nodeList.sort((o1, o2) -> -o1.compareTo(o2));
        // System.out.println(nodeList);

        int numUsed = 0;
        int totalVal = 0;
        for (int i = 0; numUsed < numWanted && i < nodeList.size(); i++) {
            Node thisNode = nodeList.get(i);
            int remainLimit = labelLimit.get(thisNode.getLabel());
            if (remainLimit > 0) {
                labelLimit.put(thisNode.getLabel(), remainLimit - 1);
                totalVal += thisNode.getVal();
                numUsed++;
            } else {
                continue;
            }
        }

        return totalVal;
    }
}

class Node implements Comparable {
    private int val;
    private int label;

    public Node(int val, int label) {
        this.val = val;
        this.label = label;
    }

    public int getVal() {
        return val;
    }

    public int getLabel() {
        return label;
    }

    @Override
    public int compareTo(Object o) {
        if (!(o instanceof Node)) {
            throw new ClassCastException();
        }
        Node anotherNode = (Node) o;
        return this.val - anotherNode.getVal();
    }

    @Override
    public String toString() {
        return String.format("[%d, %d]", val, label);
    }
}
// @lc code=end

// Accepted
// 81/81 cases passed (21 ms)
// Your runtime beats 77.47 % of java submissions
// Your memory usage beats 81.87 % of java submissions (44.7 MB)